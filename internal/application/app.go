package app

import (
	"context"
	"fmt"
	configs "task-system/cmd/config"
	"task-system/internal/infrastructure/web"
	mysql_pkg "task-system/pkg/mysql"
)

func NewApplication() *web.Server {
	ctx := context.Background()

	mysqlConnectionInput := mysql_pkg.MySqlInput{
		User:     configs.MySqlCfg.User,
		Password: configs.MySqlCfg.Password,
		Host:     configs.MySqlCfg.Host,
		Port:     configs.MySqlCfg.Port,
		Name:     configs.MySqlCfg.Name,
	}

	db, err := mysql_pkg.ConnectToDatabase(mysqlConnectionInput)

	if err != nil {
		fmt.Println(err, "Failed to connect to database")
	}

	fmt.Println(db, "Failed to connect to database")

	srv := web.NewServer(ctx)

	return srv
}
