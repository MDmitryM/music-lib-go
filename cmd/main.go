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

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("cant read cfg!")
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Can't load .env file, %s", err.Error())
	}

	app := fiber.New(fiber.Config{
		//Prefork:       true,             // включаем предварительное форкование для увеличения производительности на многоядерных процессорах
		ServerHeader:  "Fiber",          // добавляем заголовок для идентификации сервера
		ReadTimeout:   5 * time.Second,  // Тайм-ауты для чтения
		WriteTimeout:  10 * time.Second, // Тайм-ауты для записи
		CaseSensitive: true,             // включаем чувствительность к регистру в URL
		StrictRouting: true,             // включаем строгую маршрутизацию
	})

	db, err := repository.NewPostgresDB(repository.PostgresConfig{
		Host:     viper.GetString("dev_db.host"),
		Port:     viper.GetString("dev_db.port"),
		Username: viper.GetString("dev_db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("dev_db.sslmode"),
		DBName:   viper.GetString("dev_db.dbname"),
	})
	if err != nil {
		logrus.Fatalf("Connection to db failed, %s", err.Error())
	}

	if err = repository.MigratePostgres(db); err != nil {
		logrus.Fatalf("Migration failed, %s", err.Error())
	}

	repository := repository.NewRepository(db)
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
