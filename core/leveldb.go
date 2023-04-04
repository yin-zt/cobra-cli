package core

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/yin-zt/cobra-cli/utils"
	"log"
	"os"
)

var (
	db   *leveldb.DB
	err  error
	data []byte
	obj  interface{}
	opts = &opt.Options{
		CompactionTableSize: 1024 * 1024 * 20,
		WriteBuffer:         1024 * 1024 * 20,
	}
)

func (this *Cli) LevelOperate(action, key, value, path string) {
	if path == "" {
		if v, err := utils.Home(); err == nil {
			path = v
		}
	}
	if !utils.IsExist(path) {
		err = os.Mkdir(path, 0777)
		if err != nil {
			corelog.Errorf("fail to create dir %s", err)
			log.Fatalln(err)
		}
	}
	filename := path + "/" + "o.db"
	db, err = leveldb.OpenFile(filename, opts)
	if err != nil {
		corelog.Error(err)
		log.Fatalln(err)
		return
	}
	switch action {
	case "get":
		data, err = db.Get([]byte(key), nil)
		if err != nil {
			corelog.Error(err)
			log.Fatalln(err)
		}
		fmt.Println(string(data))
	case "put":
		err = db.Put([]byte(key), []byte(value), nil)
		if err != nil {
			corelog.Error(err)
			log.Fatalln(err)
		}
		fmt.Println("ok")
	}
}
