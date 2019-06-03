package main

import (

	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

func InitialMigration() {
	//Variables required for setup
	/*
	user= (using default user for postgres database)
	dbname= (using default database that comes with postgres)
	password = (password used during initial setup)
	host = (IP Address of server)
	sslmode = (must be set to disabled unless using SSL. This is not covered during tutorial)
	*/

	//DO NOT SAVE PASSWORD AS TEXT IN A PRODUCTION ENVIRONMENT. TRY USING AN ENVIRONMENT VARIABLE
	connStr := "user=postgres dbname=postgres password=123456go host=127.0.0.1 sslmode=disable"
	//driver name part of "github.com/lib/pq"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print(err)
	}
	//check postgres to see if table exists
	var checkDatabase string
	db.QueryRow("SELECT to_regclass('public.events')").Scan(&checkDatabase)
	if err != nil {
		fmt.Print(err)
	}
	//if table dose not exist then create one to use for this example
	if checkDatabase == ""{
		fmt.Println("Database Created")
		createSQL := "CREATE TABLE public.events (pk SERIAL PRIMARY KEY,count numeric,type character varying,timestamp numeric );"
		db.Query(createSQL)
	} else {
		fmt.Println("The database is already created")
	}

}