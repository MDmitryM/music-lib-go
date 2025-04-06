package repository

import (
	"fmt"
	"time"

	"github.com/MDmitryM/music-lib-go/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	maxRetries    = 10
	retryInterval = 3 * time.Second
)

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	SSLMode  string
	DBName   string
}

func NewPostgresDB(cfg PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)

	logrus.Printf("Attempting to connect to database with DSN: host=%s, port=%s, user=%s, dbname=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName)

	var db *gorm.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		logrus.Printf("Connection attempt %d/%d", i+1, maxRetries)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			logrus.Println("Database connected successfully")
			return db, nil
		}

		logrus.Printf("Failed to connect: %v. Retrying in %v...", err, retryInterval)
		time.Sleep(retryInterval)
	}

	return nil, fmt.Errorf("failed to connect after %d attempts: %w", maxRetries, err)
}

func MigratePostgres(db *gorm.DB) error {
	return db.AutoMigrate(&models.UserModel{}, &models.SongModel{})
}
