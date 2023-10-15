package testdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var (
	dbUser   = "root"
	dbPassWd = "root"
	dbAddr   = "localhost:3306"
	dbName   = "test"
)

func InitDB() (*sql.DB, error) {
	dbConfig := mysql.Config{
		User:      dbUser,
		Passwd:    dbPassWd,
		Net:       "tcp",
		Addr:      dbAddr,
		DBName:    dbName,
		ParseTime: true,
	}

	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open mysql")
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, fmt.Errorf("database is not healthy")
	}

	return db, nil
}

func SetUp(db *sql.DB, fileName string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("failed to read sql file")
	}

	_, err = db.Exec(string(file))
	if err != nil {
		log.Fatal("failed to exec sql file")
	}
}
