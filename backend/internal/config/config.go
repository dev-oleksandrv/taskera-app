package config

import (
	"github.com/spf13/viper"
	"log"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type Config struct {
	GinMode   string
	JWTSecret string
	DBConfig  DBConfig
}

var cfg *Config

func Init() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("dotenv")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	cfg = &Config{
		GinMode:   viper.GetString("GIN_MODE"),
		JWTSecret: viper.GetString("JWT_SECRET"),
		DBConfig: DBConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			Username: viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Database: viper.GetString("DB_NAME"),
		},
	}

	log.Println("Config was successfully loaded")
}

func GetConfig() *Config {
	return cfg
}
