package main

//last time we made a web server with go that displayed a simple message lets do that again .

import (
	"net/http"
	//we need to import  couple of packages to connect our database  to our server.
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	//this is the package that allows us to create a web server in go
)

func main() {
	//now to open a database connection lets see what is done.
	db, err := //from the package database
		sql.Open("sqlite3", "./database.db")
	//this pecifies the database file you are using
	if err != nil {
		panic(err)
	} //if there is an error the program will return the error.
	defer db.Close()
	//this closes the database but it does that at the end of execution which means this is executed last.
	fmt.Println("database connection established")

	//now the program should tell us that the database connection was successful and we can now create a table in the database.

	createTable := `CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL
		)`
	//CREATE TABLE IF NOT EXISTS users this tells the database to create a table if non exists and the table name is users.
	//id INTEGER PRIMARY KEY AUTOINCREMENT, THIS STATES THE FIRST FIELD IN THE USERS TABLE IS ID IT IS AN ITERGER IT IS ALSO A PRIMARY KEY AN IT WILL AUTO INCRIMENT WHEN A NEW RECORD IS ADDED TO THE TABLE.
	//the format for the other fileds is similar

	_, err = db.Exec(createTable)
	//here we are executing the create table statement and checking for errors.
	if err != nil {
		panic(err)

	}
	fmt.Println("table created successfully")
	//now when we run the program it should create a database file and a table in the database file.
	http.HandleFunc("/", home)
	//this function tells the server to listen for requests to the root URL and call the home function when a request is received.
	http.ListenAndServe(":8080", nil)
	//this function starts the server and listens for incoming requests on port 8080. the second argument is nil because we are not using any custom handlers.
}
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello Peter!"))
	//this function writes the response to the client. it takes two arguments, the first is the response writer and the second is the request. we are writing a simple message "hello world" to the response writer.
}

//lets try to add a database to our server and create a table in the database
//first we will create a datbase file with the extension .db
