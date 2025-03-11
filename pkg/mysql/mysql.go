package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser     = "user"
	dbPassword = "password"
	dbHost     = "127.0.0.1"
	dbPort     = "3306"
	dbName     = "mydatabase"
)

func ConnectToDatabase() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			fmt.Println("Conectado ao MySQL com sucesso!")
			return db, nil
		}
		fmt.Println("Tentando conectar ao MySQL...")
		time.Sleep(1 * time.Second)
	}

	return nil, fmt.Errorf("Falha ao conectar ao MySQL: %v", err)
}
