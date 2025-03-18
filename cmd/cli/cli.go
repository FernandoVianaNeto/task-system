package cli

import (
	"context"
	"fmt"
	"os"
	configs "task-system/cmd/config"
	app "task-system/internal/application"
	workers "task-system/internal/infrastructure/worker"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(taskConsumerCmd)
}

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "root - main command application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Initialize Http Server",
	Run: func(cmd *cobra.Command, args []string) {
		port := configs.ApplicationCfg.AppPort
		if port == 0 {
			os.Exit(1)
		}

		srv := app.NewApplication()
		if err := srv.Start(fmt.Sprintf(":%d", port)); err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing application: %v\n", err)
			os.Exit(1)
		}
	},
}

var taskConsumerCmd = &cobra.Command{
	Use:   "task-consumer",
	Short: "Start Task Consumer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing Task Consumer")
		configs.InitializeWorkerConfig()

		ctx := context.Background()

		workers.StartTaskStatusUpdatedConsumer(ctx)
	},
}
