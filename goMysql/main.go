package main

import "fmt"
import "database/sql"
import _"github.com/go-sql-driver/mysql"

type student struct {
	id string
	name string 
	age int
	grade int
}

func connect() (*sql.DB, error){
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err  != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery(){
	db, err := connect()
	if err !=nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err

}