package product

import (
	"fmt"
	envConfig "github.com/pius-microservices/piopos-database/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	envCfg := envConfig.LoadConfig()

	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s ", envCfg.ProductDBHost, envCfg.ProductDBPort, envCfg.ProductDBUsername, envCfg.ProductDBPassword, envCfg.ProductDBName)

	gormDb, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return gormDb, nil
}
