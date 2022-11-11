package utils

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var mutex sync.Mutex

// Gorm 实例
var DB *gorm.DB

func init() {
	DB = mySQL()
}

func mySQL() *gorm.DB {
	if db == nil {
		mutex.Lock()
		defer mutex.Unlock()

		// 配置数据库
		db_user := "root"
		db_pass := "New4you!"
		db_host := "127.0.0.1:3306"
		db_name := "beego_api_demo"

		if db_user == "" || db_pass == "" || db_name == "" {
			panic("please check db info")
		}
		dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass, db_host, db_name)
		d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		db = d
	}
	return db
}
