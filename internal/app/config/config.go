package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath("../config")
	v.SetConfigName("config")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("ошибка при инициализации конфигурационного файла: %s", err.Error())
	}

	// Завершение приложения при ошибке в загрузке переменных окружения
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("ошибка при загрузке переменных окружения: %s", err.Error())
	}

	return v, nil
}
