package dbConfig

import (
	"database/sql"
	"fmt"

	"github.com/golang/glog"
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/redis/go-redis/v9"
	//gorm_mysql "gorm.io/driver/mysql"
	//"gorm.io/gorm"
	//"gorm.io/gorm/logger"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var Envmode EnvType

func InitDB(dbname string, username string, password string, host string, port int) {
	// read credentails from config file
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, username, password, dbname)
	// only initialise connection if its not initialised before
	if DB == nil {
		Db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
		//defer Db.Close()
		err = Db.Ping()
		if err != nil {
			panic(err)
		}
		DB = Db
		glog.V(2).Infoln("Database sucessfully configured")
		return
	}
	glog.V(2).Infoln("Database already configured")

	/*
		// testing if psql connection is working fine or not
		rows, err := Db.Query("SELECT version()")
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			var result string
			err = rows.Scan(&result)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Version: %s\n", result)
		}
		fmt.Println("Established a successful connection!")
	*/

}

func SetEnvMode(mode EnvType) {
	Envmode = mode
}

func GetEnvMode() EnvType {
	return Envmode
}
