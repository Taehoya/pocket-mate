package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func InitDB(config *config) (*sql.DB, error) {
	dbConfig := mysql.Config{
		User:      config.DbUser,
		Passwd:    config.DbPasswd,
		Net:       "tcp",
		Addr:      config.DbAddr,
		DBName:    config.DbName,
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
