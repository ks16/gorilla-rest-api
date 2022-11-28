package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Debug     bool
	Name      string
	WebServer *WebServer
}

type WebServer struct {
	Port         int
	WriteTimeout uint
	ReadTimeout  uint
	IdleTimeout  uint
}

func NewConfig() *Config {
	viper.AutomaticEnv()
	setDefaults()

	_ = viper.ReadInConfig()
	webServer := &WebServer{
		Port:         viper.GetInt("PORT"),
		WriteTimeout: viper.GetUint("SERVER_WRITE_TIMEOUT"),
		ReadTimeout:  viper.GetUint("SERVER_READ_TIMEOUT"),
		IdleTimeout:  viper.GetUint("SERVER_IDLE_TIMEOUT"),
	}

	return &Config{
		Name:      viper.GetString("NAME"),
		Debug:     viper.GetBool("DEBUG"),
		WebServer: webServer,
	}
}

func setDefaults() {
	viper.SetDefault("NAME", "REST API TEMPLATE")
	viper.SetDefault("DEBUG", false)
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("SERVER_WRITE_TIMEOUT", 15)
	viper.SetDefault("SERVER_READ_TIMEOUT", 15)
	viper.SetDefault("SERVER_IDLE_TIMEOUT", 60)
}
