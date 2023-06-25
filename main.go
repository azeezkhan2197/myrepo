package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() *sql.DB {
	const (
		username = "root"
		password = "password"
		hostname = "127.0.0.1:3306"
		dbname   = "mydb"
	)

	db, err := sql.Open("mysql", "root:password@/mydb")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func main() {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM users ORDER BY id DESC")
	if err != nil {
		fmt.Println("error while selecting the users %v", err)
		panic(err.Error())
	}
	usersList := []Employee{}
	for selDB.Next() {
		user := Employee{}
		err := selDB.Scan(&user.id, &user.name)
		if err != nil {
			fmt.Println("error while scanning %v", err)
		}
		usersList = append(usersList, user)
	}
	fmt.Println(usersList)
	defer db.Close()
}

type Employee struct {
	id   int
	name string
}
