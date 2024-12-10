package database

import (
	"database/sql" // Package for SQL database interactions
	"fmt"
	"log"
	"signup-login/model"

	_ "modernc.org/sqlite" // SQLite driver
)

// Database is an interface for database operations.
type Database interface {
	GetUser(username string) (model.User, error)
	AddUser(name, username, password string) error
}

// DB is a global variable for the SQLite database connection
var DB *sql.DB

// initDB initializes the SQLite database and creates the todos table if it doesn't exist
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./app.db") // Open a connection to the SQLite database file named app.db
	if err != nil {
		log.Fatal(err) // Log an error and stop the program if the database can't be opened
	}

	// SQL statement to create the todos table if it doesn't exist
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS user (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(255) NOT NULL UNIQUE,
	username VARCHAR(255) NOT NULL UNIQUE,
	password TEXT NOT NULL
	);`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt) // Log an error if table creation fails
	}
	fmt.Println("Database Ready")
}

func CloseDB() {
	if DB != nil {
		fmt.Println("Database Closed")
		DB.Close()
	}
}

// function to get user from database
func GetUser(username string) (model.User, error) {
	rows, err := DB.Query("Select * from user where username == (?)", username)

	if err != nil {
		return model.User{}, err
	}

	defer rows.Close()

	userResult := model.User{}
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UserID,
			&user.Name,
			&user.Username,
			&user.Password); err != nil {
			return model.User{}, err
		}
		userResult = user
	}

	if userResult.Username == "" {
		return userResult, fmt.Errorf("user does not exist")
	}

	return userResult, nil

}

// function to add user to the database
func AddUser(name string, username string, password string) error {
	var user model.User
	user.Name = name
	user.Username = username
	user.Password = password
	user.HashPassword(user.Password)
	_, err := DB.Exec("INSERT INTO user(name, username, password) VALUES(?,?,?)", user.Name, user.Username, user.Password)
	return err
}
