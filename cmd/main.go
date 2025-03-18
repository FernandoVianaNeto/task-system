package main

import (
	"task-system/cmd/cli"
	configs "task-system/cmd/config"
)

func main() {
	configs.InitializeConfigs()

	cli.Execute()
}
