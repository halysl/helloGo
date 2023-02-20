package main

import (
	"encoding/json"
	"flag"
	"github.com/syndtr/goleveldb/leveldb"
	"io/ioutil"
	"log"
	"strings"
)

var levelDBPath = flag.String("path", "../", "leveldb path")
var key = flag.String("key", "", "")
var savePath = flag.String("savepath", "/tmp/leveldb", "")

func main() {
	flag.Parse()
	var db *leveldb.DB
	var err error
	if *levelDBPath != "../" {
		db, err = leveldb.OpenFile(*levelDBPath, nil)
		if err != nil {
			log.Fatal("open leveldb error:", err)
		}
	}
	res, err := getAll(db, *key)
	if err != nil {
		log.Fatal("get leveldb data error:", err)
	}
	data, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		log.Fatal("json marshal error:", err)
	}
	if err = ioutil.WriteFile(*savePath + *key + ".json", data, 0755); err != nil {
		log.Fatal("writefile error:", err)
	}
}

func getAll(l *leveldb.DB, keyType string) (map[string]string, error) {
	mapRlt := map[string]string{}
	iter := l.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()

		keyElems := strings.Split(string(key), "_")
		if keyType != "" {
			if len(keyElems) < 1 || keyElems[0] != keyType {
				continue
			}
		}
		//log.Info(string(key), "------->", string(value))
		//log.Info(strings.Join(keyElems[1:], "_"), "------->", string(value))
		mapRlt[strings.Join(keyElems[1:], "_")] = string(value)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		log.Fatal(err)
	}
	return mapRlt, err
}

/*
./getleveldb -path /opt/lotus/lotusstorage-sched/conf/db/ -key "url"
./getleveldb -path /opt/lotus/lotusstorage-sched/conf/db/ -key "binding"
./getleveldb -path /opt/lotus/lotusstorage-sched/conf/db/ -key "doingTask"
./getleveldb -path /opt/lotus/lotusstorage-sched/conf/db/ -key "taskResult"


*/
