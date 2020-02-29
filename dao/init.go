package dao

import (
	"time"
	// import mysql https://zhuanlan.zhihu.com/p/107188453
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var _DB *gorm.DB

// DB get db
func DB() *gorm.DB {
	return _DB
}
func init() {
	_DB = initDB()
}

func initDB() *gorm.DB {
	// In our docker dev environment use
	// db, err := gorm.Open("mysql", "go_web:go_web@tcp(database:3306)/go_web?charset=utf8&parseTime=True&loc=Local")
	mysqlConfig := os.Getenv("MYSQL_CONF")
	db, err := gorm.Open("mysql", mysqlConfig)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(time.Second * 300)
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
	log.Println("init db success!")
	return db
}
