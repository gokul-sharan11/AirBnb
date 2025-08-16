package config 

import (
	env "auth_service/config/env"
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

// Retruns a connection object to the database
func SetupDB() (*sql.DB, error) {
	config := mysql.NewConfig()

	config.User = env.GetString("DB_USER", "root")
	config.Passwd = env.GetString("DB_PASSWORD", "")
	config.Net = "tcp"
	config.Addr = env.GetString("DB_ADDR", "127.0.0.1:3306")
	config.DBName = env.GetString("DB_NAME", "auth_service")

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		fmt.Println("Error connecting to database", err)
		return nil, err
	}

	pingErr := db.Ping(); 
	if pingErr != nil {
		fmt.Println("Error pinging database", pingErr)
		return nil, pingErr
	}

	fmt.Println("Connected to database")

	return db, err
}