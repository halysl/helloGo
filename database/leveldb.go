package database

import "github.com/syndtr/goleveldb/leveldb"

var LevelDB *leveldb.DB

func init() {
	LevelDB, _ = leveldb.OpenFile("conf/hello.db", nil)
}
