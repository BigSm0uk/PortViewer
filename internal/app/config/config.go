package config

import (
	"errors"
	"flag"

	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	Server  ServerConfig  `yaml:"server"`
	Logging LoggingConfig `yaml:"logging"`
}

func ReadConfig() (*AppConfig, error) {
	path := flag.String("config", "", "path to config file")
	flag.Parse()
	var cfg AppConfig
	if path == nil || *path == "" {
		return nil, errors.New("invalid path to config file")
	}
	err := cleanenv.ReadConfig(*path, &cfg)
	return &cfg, err
}

type ServerConfig struct {
	Host        string `yaml:"host" env-default:"localhost"`
	Port        string `yaml:"port" env-default:"8080"`
	OpenBrowser bool   `yaml:"openBrowser" env-default:"false"`
}

type LoggingConfig struct {
	Level  string `yaml:"level" env-default:"info"`
	Format string `yaml:"format" env-default:"text"`
}
