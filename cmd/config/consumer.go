package configs

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	defaultTimeOut    time.Duration = 5
	defaultBatchLimit int           = 20
)

type WorkerConfig struct {
	Topic string `json:"topic" validate:"required"`
}

func ParseWorkerConfig(dat []byte) (*WorkerConfig, error) {
	var workerCfg *WorkerConfig

	if err := json.Unmarshal(dat, &workerCfg); err != nil {
		return nil, err
	}

	if err := validator.New().Struct(workerCfg); err != nil {
		return nil, err
	}

	return workerCfg, nil
}
