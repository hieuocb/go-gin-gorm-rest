package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/caarlos0/env"
	"github.com/hieuocb/go-gin-gorm-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Connect(dbUrl string) (*gorm.DB, error) {
	//dbUrl := "postgres://postgres:postgres@localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Warn),
	})
	if err != nil {
		return nil, err
	}
	errAM := db.AutoMigrate(&models.User{})
	if errAM != nil {
		return nil, err
	}

	DB = db
	return db, nil
}

type Schema struct {
	HttpPort    string `env:"http_port"`
	DatabaseURI string `env:"database_url"`
}

var (
	cfg Schema
)

func LoadConfig() *Schema {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)

	err := godotenv.Load(filepath.Join(currentDir, "config.yaml"))

	if err != nil {
		log.Printf("Error on load configuration file, error: %v", err)
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error on parsing configuration file, error: %v", err)
	}

	return &cfg
}

func GetConfig() *Schema {
	return &cfg
}
