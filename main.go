package main

import (
	"net/http"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	
)

func main() {
	db, err := 
		sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	} 
	defer db.Close()
	fmt.Println("database connection established")

	createTable := `CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL
		)`
	

	_, err = db.Exec(createTable)
	
	if err != nil {
		panic(err)

	}
	fmt.Println("table created successfully")
	
	http.HandleFunc("/", home)
	
	http.ListenAndServe(":8080", nil)
	
}
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello Peter!"))
	
}


