package main

import (
	"github.com/MDmitryM/music-lib-go/pkg/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("cant read cfg!")
	}

	app := handler.SetupRouts()
	app.Listen(viper.GetString("port"))
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
