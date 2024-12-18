package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MDmitryM/music-lib-go/pkg/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("cant read cfg!")
	}

	app := fiber.New(fiber.Config{
		//Prefork:       true,             // включаем предварительное форкование для увеличения производительности на многоядерных процессорах
		ServerHeader:  "Fiber",          // добавляем заголовок для идентификации сервера
		ReadTimeout:   5 * time.Second,  // Тайм-ауты для чтения
		WriteTimeout:  10 * time.Second, // Тайм-ауты для записи
		CaseSensitive: true,             // включаем чувствительность к регистру в URL
		StrictRouting: true,             // включаем строгую маршрутизацию
	})
	h := new(handler.Handler)
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
