package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const (
	HOST = "golang-mysql-1"
	PORT = "3306"
)

var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func Initialize(username, password, database string) (Database, error) {

	log.Println("Connecting to db")
	// conn, err := sql.Open("mysql", "user:password@tcp(golang-mysql-1:3306)/database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, HOST, PORT, database)

	db := Database{}
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}
