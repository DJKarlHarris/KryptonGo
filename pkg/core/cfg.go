package core

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerConfig ServerConfig `yaml:"ServerConfig"`
}

type ServerConfig struct {
	Test  int      `yaml:"Test"`
	Table []string `yaml:"Table"`
}

func (app *App) LoadConfig(path string) (err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		SLOG().Errorf("read file fail %s", err.Error())
		return err
	}

	if err = yaml.Unmarshal(data, &app.config); err != nil {
		SLOG().Errorf("unmarshal fail %s", err.Error())
		return err
	}

	return err
}

func GetConfig() *Config {
	return &app.config
}
