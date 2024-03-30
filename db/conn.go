package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var onceMySQLInstance sync.Once
var mySQLInstance *sql.DB

var onceMySQLAuthDBInstance sync.Once
var AuthDBInstance *sql.DB

func GetMySQLInstanceForAUTH() *sql.DB {
	if AuthDBInstance == nil {
		onceMySQLAuthDBInstance.Do(setUpMySQLConnection)
	}
	return AuthDBInstance
}

func GetMySQLInstance() *sql.DB {
	if mySQLInstance == nil {
		onceMySQLInstance.Do(setUpMySQLConnection)
	}
	return mySQLInstance
}

func setUpMySQLConnection() {
	var err error

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "root1234", "localhost", "3306", "astra_db")
	mySQLInstance, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Printf("could not connect to mysql database  : %s", err)
		return
	}
}
