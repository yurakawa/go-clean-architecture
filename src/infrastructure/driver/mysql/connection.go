package mysql

import (
	"cleanarchitecture/adapter/gateway"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type config struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
	Options  string
}


func Connect() *gorm.DB {
	var err error

	dbConfig := config {
		Username: "root",
		Password: "root",
		Host: "127.0.0.1",
		Port: 3306,
		Database: "hoge",
		Options: "parseTime=true&loc=Asia%2FTokyo",
	}
	db, err = gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.Options))

	if err != nil {
		panic(err)
	}
	if !db.HasTable(&gateway.User{}) {
		if err := db.Table("users").CreateTable(&gateway.User{}).Error; err != nil {
			panic(err)
		}
	}
	return db
}

func CloseConn() {
	db.Close()
}
