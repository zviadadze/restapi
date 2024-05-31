package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/zviadadze/userver/api/handler"
)

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func loadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func StartServer(configPath string) error {
	config, err := loadConfig(configPath)
	if err != nil {
		return err
	}

	mux := handler.RegisterRoutes()

	if err := http.ListenAndServe(config.String(), mux); err != nil {
		return err
	}

	return nil
}

func (c *Config) String() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)
}
