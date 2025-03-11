package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ApplicationCfg *ApplicationConfig
	MySqlCfg       *MySqlConfig
)

type ApplicationConfig struct {
	AppPort int
}

type MySqlConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

// type KafkaConfig struct {
// 	BrokersHost           string
// 	PushNotificationTopic string
// } // TODO LATER

func initialize() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func InitializeConfigs() {
	initialize()
	initializeApplicationConfigs()
	initializeMySqlConfings()
}

type Sentry struct {
	DSN        string
	Debug      bool
	SampleRate float64
}

func getEnv(key string, defaultVal string) string {
	value, exists := os.LookupEnv(key)

	if exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")

	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func initializeApplicationConfigs() {
	if ApplicationCfg == nil {
		ApplicationCfg = &ApplicationConfig{
			AppPort: getEnvAsInt("APP_PORT", 80),
		}
	}
}

func initializeMySqlConfings() {
	if MySqlCfg == nil {
		MySqlCfg = &MySqlConfig{
			User:     getEnv("MYSQL_USER", ""),
			Password: getEnv("MYSQL_PASSWORD", ""),
			Host:     getEnv("MYSQL_HOST", ""),
			Port:     getEnv("MYSQL_PORT", ""),
			Name:     getEnv("MYSQL_NAME", ""),
		}
	}
}

// func initializeKafkaConfigs() {
// 	if KafkaCfg == nil {
// 		KafkaCfg = &KafkaConfig{
// 			BrokersHost: getEnv("KAFKA_BROKER_HOSTS", ""),
// 		}
// 	}
// } // TODO LATER
