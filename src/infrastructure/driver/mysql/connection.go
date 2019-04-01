package mysql

import (
	"cleanarchitecture/adapter/gateway"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var dbConn *gorm.DB

func Connect() *gorm.DB {
	var err error

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetInt(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	Options := "parseTime=true&loc=Asia%2FTokyo"

	dbConn, err = gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
			dbUser,
			dbPass,
			dbHost,
			dbPort,
			dbName,
			Options,
		))

	if err != nil {
		panic(err)
	}
	if !dbConn.HasTable(&gateway.User{}) {
		if err := dbConn.Table("users").CreateTable(&gateway.User{}).Error; err != nil {
			panic(err)
		}
	}
	return dbConn
}

func CloseConn() {
	dbConn.Close()
}
