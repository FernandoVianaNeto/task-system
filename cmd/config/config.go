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
	AppPort            int
	PasswordSecretHash string
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
			AppPort:            getEnvAsInt("APP_PORT", 80),
			PasswordSecretHash: getEnv("PASSWORD_SECRET_HASH", ""),
		}
	}
}

func initializeMySqlConfings() {
	if MySqlCfg == nil {
		MySqlCfg = &MySqlConfig{
			User:     getEnv("MYSQL_USER", "user"),
			Password: getEnv("MYSQL_PASSWORD", "password"),
			Host:     getEnv("MYSQL_HOST", "127.0.0.1"),
			Port:     getEnv("MYSQL_PORT", "3307"),
			Name:     getEnv("MYSQL_NAME", "mydatabase"),
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
