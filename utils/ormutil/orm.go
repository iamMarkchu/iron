package ormutil

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var engine *xorm.Engine
var err error

func init() {
	engine, err = xorm.NewEngine("mysql", "root:root@/training_base?charset=utf8")
	if err != nil {
		log.Fatalf("init xorm faild%s", err.Error())
	}
	engine.ShowSQL(true)
}

func GetInstance() *xorm.Engine {
	return engine
}
