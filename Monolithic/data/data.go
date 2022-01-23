package data

import (
	"database/sql"
	"fmt"
	"log"
	"monolithicdemo/models"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Create a connection to the database in MySQL
func getConnection() *sql.DB {
	var db *sql.DB
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "messages",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func GetMessages() (models.Messages, error) {

	db := getConnection()
	defer db.Close()
	messages := models.Messages{}

	rows, err := db.Query("CALL GetMessages();")
	if err != nil {
		return messages, fmt.Errorf("message: %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		msg := models.Message{}

		if err := rows.Scan(&msg.Id, &msg.Message); err != nil {
			return messages, fmt.Errorf("message: %v", err)
		}
		messages.List = append(messages.List, msg)
	}

	if err := rows.Err(); err != nil {
		return messages, fmt.Errorf("message: %v", err)
	}

	return messages, nil
}

/*
func GetUserMessages() models.UserMessages {

}

func AddMessage() {
	// TODO: Add to Database
	fmt.Println("Adding data...")
}

func AddUserMessage() {

}

func UpdateUserMessage() {

}

func DeleteUserMessage() {
	// TODO: Delete from Database
	fmt.Println("Deleting data...")
}

*/
