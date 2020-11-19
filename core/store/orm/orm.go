package orm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	cfg "github.com/spf13/viper"
	"log"
)

var (
	db       *gorm.DB
	dsn      string
	logModel bool
	err      error
)

// 仅支持一个db实例
// todo 扩展到多db实例
func Init() {
	dsn = cfg.GetString("mysql.dsn")
	logModel = cfg.GetBool("mysql.logMode")
	if len(dsn) == 0 {
		fmt.Printf("dsn can not be empty \n")
	}
	if db, err = gorm.Open("mysql", dsn); err != nil {
		log.Panicf("gorm Open err: %s\n", err)
	}
	db.LogMode(logModel)
	fmt.Printf("ping:%s\n", db.DB().Ping())
}

func GetInstance() *gorm.DB {
	return db
}
