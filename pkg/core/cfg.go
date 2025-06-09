package core

import (
	"os"

	"gopkg.in/yaml.v3"
)

// tips首字母大写保证字段可见
type Config struct {
	ServerConfig ServerConfig `yaml:"ServerConfig"`
}

type Db struct {
	Addr   string `yaml:"Addr"`
	Passwd string `yaml:"Passwd"`
	Port   string `yaml:"Port"`
}

type ServerConfig struct {
	Db Db `yaml:"Db"`
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
