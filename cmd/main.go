package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MDmitryM/music-lib-go/pkg/handler"
	"github.com/MDmitryM/music-lib-go/pkg/repository"
	"github.com/MDmitryM/music-lib-go/pkg/service"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title						Music lib API
// @version					    1.0
// @description				    API server for music lib applicatiom
// @host						localhost:8080
// @BasePath					/
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("cant read cfg!")
	}

	var conf repository.PostgresConfig
	var redisConf repository.RedisConfig
	env := os.Getenv("ENV")
	logrus.Printf("env = %s\n", env)
	if env != "production" {
		//Открываем .env только в dev
		if err := godotenv.Load(); err != nil {
			logrus.Fatalf("Can't load .env file, %s", err.Error())
		}
		conf = repository.PostgresConfig{
			Host:     viper.GetString("dev_db.host"),
			Port:     viper.GetString("dev_db.port"),
			Username: viper.GetString("dev_db.username"),
			Password: os.Getenv("DB_PASSWORD"),
			SSLMode:  viper.GetString("dev_db.sslmode"),
			DBName:   viper.GetString("dev_db.dbname"),
		}
		redisConf = repository.RedisConfig{
			Host:     viper.GetString("redis_dev.host"),
			Port:     viper.GetString("redis_dev.port"),
			Password: os.Getenv("REDIS_DB_PASSWORD"),
			DB:       viper.GetInt("redis_dev.DB"),
		}
	} else {
		conf = repository.PostgresConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			Password: os.Getenv("DB_PASSWORD"),
			SSLMode:  viper.GetString("db.sslmode"),
			DBName:   viper.GetString("db.dbname"),
		}
		redisConf = repository.RedisConfig{
			Host:     viper.GetString("redis.host"),
			Port:     viper.GetString("redis.port"),
			Password: os.Getenv("REDIS_DB_PASSWORD"),
			DB:       viper.GetInt("redis.DB"),
		}
	}

	app := fiber.New(fiber.Config{
		//Prefork:       true,             // включаем предварительное форкование для увеличения производительности на многоядерных процессорах
		ServerHeader:  "Fiber",          // добавляем заголовок для идентификации сервера
		ReadTimeout:   5 * time.Second,  // Тайм-ауты для чтения
		WriteTimeout:  10 * time.Second, // Тайм-ауты для записи
		CaseSensitive: true,             // включаем чувствительность к регистру в URL
		StrictRouting: true,             // включаем строгую маршрутизацию
	})

	db, err := repository.NewPostgresDB(conf)
	if err != nil {
		logrus.Fatalf("Connection to db failed, %s", err.Error())
	}

	if err = repository.MigratePostgres(db); err != nil {
		logrus.Fatalf("Migration failed, %s", err.Error())
	}

	redisDB, err := repository.NewRedisDB(redisConf)
	if err != nil {
		logrus.Fatalf("redis error, %s", err.Error())
	}

	repository := repository.NewRepository(db, redisDB)
	service := service.NewService(repository)

	h := handler.NewHandler(service)
	h.SetupRouts(app)

	go func() {
		if err := app.Listen(viper.GetString("port")); err != nil {
			logrus.Fatalf("server stopped, %s", err.Error())
		}
	}()
	logrus.Println("Server started on port ", viper.GetString("port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := app.Shutdown(); err != nil {
		logrus.Fatalf("Error during shutdown: %v", err)
	}

	logrus.Println("Server stopped")
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
