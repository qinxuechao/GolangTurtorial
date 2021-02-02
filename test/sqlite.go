package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, err := sql.Open("sqlite3", "./bogo.db")
	if err != nil {
		fmt.Printf("open sql error %s ", err)
	}
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER primary key, firstname TEXT, lastname TEXT)")
	statement.Exec()

	statement, _ = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?) ")
	statement.Exec("Lucas", "Qin")

	rows, _ := database.Query("SELECT id, firstname, lastname from people")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + " : " + firstname + " " + lastname)
	}
}
