package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgres(config *AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.DBAddress, config.DBUsername, config.DBPassword, config.DBName, config.DBPort,
	)

	database, er := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if er != nil {
		log.Fatalf("failed to open database %v", er)
	}

	db, err := database.DB()

	if err != nil {
		log.Fatalf("failed getting database %v", err)
	}

	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	return database
}
