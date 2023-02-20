package database

import (
"database/sql"
"fmt"
_ "github.com/go-sql-driver/mysql"
logging "github.com/ipfs/go-log/v2"
"github.com/spf13/viper"
"golang.org/x/xerrors"
"os"
"strings"
"syscall"
)

const configPath string = "CONFIG_PATH"
const TableStoragePath = "t_storage_path"
const TableSectorStorage = "t_sector_storage"

var DBViper = viper.New()
var DB *sql.DB
var log = logging.Logger("database")

func init() {
	if err := initDBViperConfig(); err != nil {
		panic(err)
	}
	if err := initDB(); err != nil {
		panic(err)
	}
}

func initDBViperConfig() error {
	cfgPath := os.Getenv(configPath)
	if cfgPath == "" {
		return xerrors.Errorf("配置文件路径未指定，请通过%s环境变量指定配置文件的路径！", configPath)
	}

	DBViper.SetConfigType("toml")
	DBViper.SetConfigFile(cfgPath)
	if err := DBViper.ReadInConfig(); err != nil {
		log.Warnf("read config err %+v,try to reset config file with default value", err)
		return err
	}
	return nil
}

func initDB() error {
	user := DBViper.GetString("db.user")
	password := DBViper.GetString("db.password")
	ip := DBViper.GetString("db.ip")
	port := DBViper.GetInt("db.port")
	database := DBViper.GetString("db.database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, ip, port, database)
	log.Infof("dsn:%s", dsn)
	DB, _ = sql.Open("mysql", dsn)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		log.Error("opon database fail!!!", err)
		syscall.Exit(-1)
	}
	log.Info("connnect success!\n")
	return nil
}

func InsertOne(statement string, table string, args ...interface{}) error {
	statement = strings.Replace(statement, "table", table, 1)
	log.Infof("sql statement:%s", statement)
	stmt, err := DB.Prepare(statement)
	if err != nil {
		log.Errorf("db Prepare error: %s", err)
		return err
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		log.Errorf("db stmt exec error: %s", err)
		return err
	} else {
		lineAffected, err1 := res.RowsAffected()
		if err1 != nil || lineAffected != 1 {
			log.Errorf("db rows affected error: %s", err)
			return err1
		} else {
			log.Infof("insert into db success, args:%+v", args)
		}
	}
	return nil
}

func getStoragePathMap() (map[int]string, error) {
	var (
		_id int
		_path string
	)
	res := make(map[int]string)
	stmt, err := DB.Prepare(fmt.Sprintf("select id,storage_path from %s", TableStoragePath))
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&_id, &_path)
		if err != nil {
			continue
		}
		res[_id] = _path
	}
	if err := rows.Close(); err != nil {
		return res, err
	}
	return res, nil
}

func GetStoragePathID(p string) (int, error) {
	_id := 0
	mapStoragePath, err := getStoragePathMap()
	if err != nil {
		log.Error("Get storage path err:", err)
		return _id, err
	}
	pAddSuffix := p + "/"
	for k, v := range mapStoragePath {
		if strings.HasPrefix(pAddSuffix, v) {
			_id = k
		}
	}
	return _id, nil
}
