package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
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
	//lets create a password variable an d hash it using bcrypt
	password := "Password123"
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	//lets run this and see how our password apears now
	//lets try to convert it to strings instead of viewing the byte version of the string
	//now to insert a user into the database we will use the following code
	insertUser := `INSERT INTO users(email, password, first_name, last_name) VALUES(?, ?, ?, ?)`
	_, err = db.Exec(insertUser, "john.doe@example.com", string(hashedPass), "John", "Doe")
	if err != nil {
		panic(err)
	}
	fmt.Println("user inserted successfully")
	//that is how we can insert a user into the database and also store the password safely using bcrypt hashing. Now we can run our server and test it out.
	http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)

}
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello Peter!"))

}

//today we will try to insert a user and also learn how to store password safely.
//first we install the bcrypt package for password hashing. using the command: go get golang.org/x/crypto/bcrypt
//then in our go fill we will import it
