package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AppName   string `yaml:"app_name"`
	Port      int    `yaml:"port"`
	LogFile   string `yaml:"log_file"`
	JWTSecret string `yaml:"jwt_secret"`
	DB        struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"db"`
}

var AppConfig Config

func LoadConfig() error {
	err := cleanenv.ReadConfig("config/config.yaml", &AppConfig)
	if err != nil {
		log.Println("Failed to load config:", err)
		return err
	}
	return nil
}
