package main

import (
	"fmt"
	"os"
	"task-system/cmd/configs"
	app "task-system/internal/application"
)

func main() {
	configs.InitializeConfigs()

	port := configs.ApplicationCfg.AppPort
	if port == 0 {
		os.Exit(1)
	}

	srv := app.NewApplication()
	if err := srv.Start(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing application: %v\n", err)
		os.Exit(1)
	}
}
