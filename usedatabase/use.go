package usedatabase

import (
	"encoding/json"
	"fmt"
	"github.com/halysl/hellogo/database"
	log "github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
	dbutil "github.com/syndtr/goleveldb/leveldb/util"
	"time"
)

const sectorIDFailedInsert = "sectorIDFailedInsert:"
var positionByStoragePath = make(map[string]int)


func DeclareSectors(p string) error {
	start := time.Now()
	defer log.Infof("scan sectors(on db) cost: %s", time.Since(start))
	var db = database.DB
	storagePathID, err := database.GetStoragePathID(p)
	if err != nil {
		return err
	}
	// 获取 sector
	sectors := make([]string, 0, 0)
	// 每次最多取500条数据，同时会更新 id
	stmt, err := db.Prepare(fmt.Sprintf(`select id,sector_id from t_sector_storage where storage_path_id=? and state=1 and id > ? limit ?`))
	if err != nil {
		log.Error("database prepare err:", err)
		return err
	}
	defer stmt.Close()
	continueFlag := true
	size := 500
	countFor := 0
	for {
		log.Infof("storage_path:%s,storage_path_id:%d,count_for:%d,start_index:%d", p, storagePathID, countFor, positionByStoragePath[p])
		countFor++
		if !continueFlag {
			break
		}
		continueFlag = false
		rows, err := stmt.Query(storagePathID, positionByStoragePath[p], size)
		if err != nil {
			return err
		}
		var (
			tableID  int
			sectorID string
		)
		countRow := 0
		for rows.Next() {
			countRow++
			continueFlag = true
			err := rows.Scan(&tableID, &sectorID)
			if err != nil {
				return err
			}
			sectors = append(sectors, sectorID)
			positionByStoragePath[p] = tableID
		}
		if err := rows.Close(); err != nil {
			log.Error("declare sector in database error:", err)
		}
		log.Infof("count_rows:%d", countRow)
	}
	return nil
}

func insertSector(args ...interface{}) error {
	return database.InsertOne(`INSERT INTO table(sector_id, storage_path_id, state, worker_name) VALUES(?,?,?,?)`, database.TableSectorStorage, args...)
}

type InsertSectorStorageData struct {
	SectorID    string
	StoragePathID int
	State       bool
	WorkerName  string
}


func UpdateSectorStoragePath() error {
	sList := []*InsertSectorStorageData{
		{
			SectorID:    "s-t0100-1025",
			StoragePathID: 1,
			State:       true,
			WorkerName:  "localhost",
		},
		{
			SectorID:    "s-t0100-1026",
			StoragePathID: 1,
			State:       true,
			WorkerName:  "localhost",
		},
		{
			SectorID:    "s-t0100-1027",
			StoragePathID: 2,
			State:       true,
			WorkerName:  "localhost",
		},
		{
			SectorID:    "s-t0100-1028",
			StoragePathID: 2,
			State:       true,
			WorkerName:  "localhost",
		},
		{
			SectorID:    "s-t0100-1029",
			StoragePathID: 2,
			State:       true,
			WorkerName:  "localhost",
		},
	}
	for _, s := range sList {
		err := insertSector(s.SectorID, s.StoragePathID, s.State, s.WorkerName)
		if err != nil {
			log.Errorf("%s insertIntoDB error: %s", s.SectorID, err)
			if err = addDbInsertFailedTask(s); err != nil {
				log.Errorf("%s addDbInsertFailedTask error: %s", s.SectorID, err)
				return err
			}
		}
	}
	return nil
}

func addDbInsertFailedTask(s *InsertSectorStorageData) error {
	key := sectorIDFailedInsert + s.SectorID
	data, _ := json.Marshal(s)
	return database.LevelDB.Put([]byte(key), data, nil)
}

func TryAsyncInsert() {
	// 每60秒遍历下db中插入数据库失败的数据
	fmt.Println("start try async")
	for {
		fmt.Println("into for...")
		list := make([]*InsertSectorStorageData, 0)
		iter := database.LevelDB.NewIterator(dbutil.BytesPrefix([]byte(sectorIDFailedInsert)), nil)
		for iter.Next() {
			key := iter.Key()
			fmt.Println("leveldb key:", string(key))
			value := iter.Value()
			var s InsertSectorStorageData
			err := json.Unmarshal(value, &s)
			fmt.Printf("unmarsh val:%+v\n", &s)
			if err != nil {
				log.Printf("tryAsyncInsert Unmarshal intodb err: %s", err)
			} else {
				list = append(list, &s)
			}
		}
		iter.Release()
		err := iter.Error()
		if err != nil {
			log.Printf("tryAsyncInsert iter err: %s", err)
		}
		okList := make([]string, 0, len(list))
		for _, li := range list {
			sectorID := li.SectorID
			log.Printf("%s tryAsyncInsert into db", sectorID)
			if err := insertSector(li.SectorID, li.StoragePathID, li.State, li.WorkerName); err == nil {
				okList = append(okList, sectorID)
			}
		}
		fmt.Println("-------------------------")
		fmt.Println(okList)
		fmt.Println("-------------------------")
		// 通知成功的删掉
		if len(okList) > 0 {
			batch := new(leveldb.Batch)
			for _, sid := range okList {
				key := sectorIDFailedInsert + sid
				batch.Delete([]byte(key))
				fmt.Println("batch...,key:", key)
			}
			err = database.LevelDB.Write(batch, nil)
			if err != nil {
				log.Printf("tryAsyncInsert db batch del err: %s", err)
			}
		}
		time.Sleep(20 * time.Second)
	}
}


func Testoffline() error {
	sList := []*InsertSectorStorageData{
		{
			SectorID:    "s-t0100-1025",
			StoragePathID: 1,
			State:       true,
			WorkerName:  "localhost",
		},
		{
			SectorID:    "s-t0100-1026",
			StoragePathID: 1,
			State:       true,
			WorkerName:  "localhost",
		},
		{
			SectorID:    "s-t0100-1027",
			StoragePathID: 2,
			State:       true,
			WorkerName:  "localhost",
		},
		{
			SectorID:    "s-t0100-1028",
			StoragePathID: 2,
			State:       true,
			WorkerName:  "localhost",
		},
		{
			SectorID:    "s-t0100-1029",
			StoragePathID: 2,
			State:       true,
			WorkerName:  "localhost",
		},
	}
	for _, s := range sList {
		//err := insertSector(s.SectorID, s.StoragePathID, s.State, s.WorkerName)
		if true {
			if err := addDbInsertFailedTask(s); err != nil {
				log.Errorf("%s addDbInsertFailedTask error: %s", s.SectorID, err)
				return err
			}
		}
	}
	return nil
}

func AddPath(p string) error {
	if err := database.InsertOne(`INSERT INTO table(storage_path) VALUES(?)`, database.TableStoragePath, p); err != nil {
		return err
	}
	return nil
}
