package mysql_pkg

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlInput struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func ConnectToDatabase(input MySqlInput) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", input.User, input.Password, input.Host, input.Port, input.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			fmt.Println("Successfully connected")
			return db, nil
		}
		fmt.Println("Trying to connect to mysql database")
		time.Sleep(1 * time.Second)
	}

	return nil, fmt.Errorf("Failed to connect to mysql database: %v", err)
}
