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

func GetMessages(userid string) (models.UserMessages, error) {
	db := getConnection()
	defer db.Close()

	messages := models.UserMessages{}

	rows, err := db.Query("CALL GetUserMessages(?);", userid)
	if err != nil {
		return messages, fmt.Errorf("message: %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		msg := models.UserMessage{}

		if err := rows.Scan(&msg.UserId, &msg.MessageId, &msg.Message); err != nil {
			return messages, fmt.Errorf("message: %v", err)
		}
		messages.List = append(messages.List, msg)
	}

	if err := rows.Err(); err != nil {
		return messages, fmt.Errorf("message: %v", err)
	}

	return messages, nil
}

func AddMessage(umsg models.UserMessage) (models.UserMessage, error) {
	db := getConnection()
	defer db.Close()
	//var lastID int

	res, err := db.Query("CALL AddUserMessage(?, ?)", umsg.UserId, umsg.Message)
	if err != nil {
		return umsg, err
	}
	defer res.Close()

	res.Next()
	res.Scan(&umsg.MessageId)

	return umsg, nil
}

/*

func AddUserMessage() {

}

func UpdateUserMessage() {

}

func DeleteUserMessage() {
		db := getConnection()
	defer db.Close()

	result, err := db.Exec("CALL AddUserMessage(?,?)", umsg.UserId, umsg.Message)
	if err != nil {
		return umsg, fmt.Errorf("addMessage: %v", err)
	}

	umsg.MessageId, err = result.LastInsertId()
	if err != nil {
		return umsg, fmt.Errorf("addMessage: %v", err)
	}
	return umsg, nil
}

*/
