package db

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ErrNotFound = fmt.Errorf("record not found")

var datetimePrecision = 2

type Connection struct {
	*gorm.DB
}

func Initialize(userName, password, HOST, PORT, databaseName string) (*Connection, error) {
	log.Println("Connecting to db")

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", userName, password, HOST, PORT, databaseName)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       DSN,                // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,                // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,               // disable datetime precision support, which not supported before MySQL 5.6
		DefaultDatetimePrecision:  &datetimePrecision, // default datetime precision
		DontSupportRenameIndex:    true,               // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,               // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,              // smart configure based on used version
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Database connection established")
	return &Connection{db}, nil
}

// _ "github.com/go-sql-driver/mysql"

// const (
// 	HOST = "golang-mysql-1"
// 	PORT = "3306"
// )
