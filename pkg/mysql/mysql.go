package mysql_pkg

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlInput struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var DB *gorm.DB

func ConnectToDatabase(input MySqlInput) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", input.User, input.Password, input.Host, input.Port, input.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to the database:", err)

		return nil, err
	}
	DB = db
	log.Println("Database successfully connected")

	return db, nil
}
