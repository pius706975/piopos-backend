package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Mode      string

	UserDBPort     string
	UserDBUsername string
	UserDBPassword string
	UserDBName     string
	UserDBHost     string

	ProductDBPort     string
	ProductDBUsername string
	ProductDBPassword string
	ProductDBName     string
	ProductDBHost     string
}

var (
	envConfig *Config
	once      sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file")
		}

		envConfig = &Config{
			Mode:      os.Getenv("MODE"),

			// User DB
			UserDBPort:     os.Getenv("USER_DB_PORT"),
			UserDBUsername: os.Getenv("USER_DB_USERNAME"),
			UserDBPassword: os.Getenv("USER_DB_PASSWORD"),
			UserDBName:     os.Getenv("USER_DB_NAME"),
			UserDBHost:     os.Getenv("USER_DB_HOST"),

			// Product DB
			ProductDBPort:     os.Getenv("PRODUCT_DB_PORT"),
			ProductDBUsername: os.Getenv("PRODUCT_DB_USERNAME"),
			ProductDBPassword: os.Getenv("PRODUCT_DB_PASSWORD"),
			ProductDBName:     os.Getenv("PRODUCT_DB_NAME"),
			ProductDBHost:     os.Getenv("PRODUCT_DB_HOST"),
		}
	})

	return envConfig
}
