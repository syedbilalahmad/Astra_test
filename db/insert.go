package db

import (
	"astra/model"
	"database/sql"
	"fmt"
	"log"
	"sync"
)

func InsertDataToDB(data model.Data) {
	mysqlInsat := GetMySQLInstance()
	wg := &sync.WaitGroup{}
	for _, v := range data.Emp {
		wg.Add(1)
		go insert(mysqlInsat, v, wg)
		wg.Wait()
	}
}

func insert(mysqlInsat *sql.DB, v model.Employee, wg *sync.WaitGroup) {
	defer wg.Done()
	// Insert data into the table
	insertStmt, err := mysqlInsat.Prepare("INSERT INTO employee (name, email , mobile) VALUES (?, ? , ?)")
	if err != nil {
		log.Printf("err %v", err.Error())
	}
	defer insertStmt.Close()

	// Execute the insert statement
	_, execerr := insertStmt.Exec(v.Name, v.Email, v.Mobile)
	if execerr != nil {
		log.Printf("err %v", err.Error())
	}
	fmt.Println("Data inserted successfully!")
}
