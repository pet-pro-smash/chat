package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Conf struct {
	HttpServer HttpServer `mapstructure:"HTTP_SERVER"`
}

type HttpServer struct {
	Addr string `mapstructure:"ADDR"`
	Port int    `mapstructure:"PORT"`
}

// Пытаеся получить файл с конфигом и распарсить
func LoadConfig() (*Conf, error) {

	// путь до каталога
	p, err := filepath.Abs("../configs/")
	if err != nil {
		return defaultConfig(), nil
	}

	// путь к файлу
	pf := filepath.Join(p, "/config.yaml")

	// проверим файл
	err = checkFile(pf)
	if err != nil {
		return defaultConfig(), nil
	}

	viper.SetConfigFile(pf)
	viper.AutomaticEnv()

	// читаем файл
	err = viper.ReadInConfig()
	if err != nil {
		return defaultConfig(), nil
	}

	//парсим файл
	config := new(Conf)
	err = viper.Unmarshal(&config)
	if err != nil {
		return defaultConfig(), nil
	}

	// сообщение в консол о конфиге
	infoConfig(config)

	// все впорядке, возвращаем конфиг из файла
	return config, nil
}

// конфиг по умолчанию (если файл не доступен)
func defaultConfig() *Conf {

	c := &Conf{
		HttpServer: HttpServer{
			Port: 8081,
		},
	}

	// сообщение в консол о конфиге
	infoConfig(c)

	return c
}

// проверим наличие кофигурационного файла
func checkFile(p string) error {
	_, err := os.Stat(p)
	if err != nil {
		return err
	}

	return nil
}

//сообщение от запуске серверов
func infoConfig(c *Conf) {
	log.Printf("сервер будет запущен по адресу http://127.0.0.1:%d\n", c.HttpServer.Port)
}
