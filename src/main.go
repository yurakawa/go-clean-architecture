package main

import (
	"cleanarchitecture/infrastructure"
	"cleanarchitecture/infrastructure/driver/mysql"
)

func main() {
	defer mysql.CloseConn()
	infrastructure.Router()
}
