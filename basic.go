package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // loading the sql drivers anonymously
)

var (
	database           string
	uid                int
	username, password string
)

func main() {
	// db, err := sql.Open("mysql", "root:qwerty@tcp(127.0.0.1:3306)/go_chat_db")
	db, err := sql.Open("mysql", "root:qwerty@tcp(127.0.0.1:3306)/")

	if err != nil {
		fmt.Println("sql.Open :", err)
	}
	defer db.Close()

	err = db.Ping()
	fmt.Println("db.Ping() :", err)

	rows, err := db.Query("show databases")

	/*
		rows, err := db.Query("select * from logins where uid = ?", 10)
		if err != nil {
			fmt.Println("db.Query() :", err)
		}*/

	defer rows.Close()

	for rows.Next() {
		// err = rows.Scan(&uid, &username, &password)
		err = rows.Scan(&database)

		if err != nil {
			fmt.Println("row.Scan() :", err)
		}

		fmt.Println(database)

	}
	// res, err := db.Exec("insert into logins(uid, username, password) values(2, 'jayesh452','passwd');")
	// if err != nil {
	// 	fmt.Println("db.Exec :", err)
	// }

	// id := res.LastInsertId()
	// fmt.Println(res.Columns())
	// n, err := res.RowsAffected()

	// fmt.Println("successful insert into", n)

}
