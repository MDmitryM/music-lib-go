package repository

import (
	"fmt"
	"os"

	"github.com/MDmitryM/music-lib-go/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	dsn := os.Getenv("DATABASE_URL")
	logrus.Printf("dsn = %s", dsn)
	//for render com
	if dsn != "" {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil
	}
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	logrus.Println("db connected successfully")
	return db, nil
}

func MigratePostgres(db *gorm.DB) error {
	return db.AutoMigrate(&models.UserModel{}, &models.SongModel{})
}
