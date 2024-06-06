package dboperation

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Connection() (*sql.DB, error) {
	fmt.Println("my sql start")
	db, err = sql.Open("mysql", "root:Abyakta@123@tcp(localhost:3306)/employee")
	if err != nil {

		return db, err
	}
	fmt.Println("my sql connection sucessfully ")
	return db, nil
}
