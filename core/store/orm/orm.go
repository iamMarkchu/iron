package orm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iamMarkchu/iron/core/config"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
	dsn string
	logModel bool
	err error
)

// 仅支持一个db实例
// todo 扩展到多db实例
func Init()  {
	dsn = config.GetInstance().GetString("mysql.dsn")
	logModel = config.GetInstance().GetBool("mysql.logMode")
	if len(dsn) == 0 {
		fmt.Printf("dsn can not be empty \n")
	}
	if db,err = gorm.Open("mysql", dsn); err != nil {
		fmt.Printf("gorm Open err: %s\n", err)
	}
	db.LogMode(logModel)
	fmt.Println("ping:", db.DB().Ping())
}

func GetInstance() *gorm.DB {
	return db
}