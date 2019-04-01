package main

import (
	"cleanarchitecture/infrastructure"
	"cleanarchitecture/infrastructure/driver/mysql"
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {


	dbConn := mysql.Connect()
	defer mysql.CloseConn()
	infrastructure.Router(dbConn)
}
