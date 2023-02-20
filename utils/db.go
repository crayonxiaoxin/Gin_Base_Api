package utils

import (
	"fmt"
	"os"
	"sync"

	"github.com/Unknwon/goconfig"
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

		// 读取环境变量（只有 dev=1 时，才是开发环境，其他默认生产环境）
		dev := os.Getenv("dev")
		isProd := dev != "1"

		// 加载配置文件
		cf, err := goconfig.LoadConfigFile("./conf.ini")
		if err != nil {
			panic("Could not load config file.")
		}

		// 根据环境变量 env 的值获取对应的配置
		var section string
		if isProd {
			section = "prod"
		} else {
			section = "dev"
		}
		db_user := cf.MustValue(section, "db_user")
		db_pass := cf.MustValue(section, "db_pass")
		db_host := cf.MustValue(section, "db_host")
		db_name := cf.MustValue(section, "db_name")
		fmt.Printf("db_user: %v\n", db_user)
		fmt.Printf("db_pass: %v\n", db_pass)
		fmt.Printf("db_host: %v\n", db_host)
		fmt.Printf("db_name: %v\n", db_name)

		// 配置数据库
		// db_user := "root"
		// db_pass := "New4you!"
		// db_host := "127.0.0.1:3306"
		// db_name := "beego_api_demo"

		// db_user := "dev"
		// db_pass := "cXxxKGTiHBGhHEiA"
		// db_host := "127.0.0.1:3306"
		// db_name := "dev"

		if db_user == "" || db_pass == "" || db_name == "" {
			panic("please check db info")
		}
		dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass, db_host, db_name)
		fmt.Printf("dsn: %v\n", dsn)
		d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		db = d
	}
	return db
}
