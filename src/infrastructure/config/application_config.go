package config

import (
	"os"
)

const envPort = "APP_PORT"

// GetApplicationConfig returns data config for api
func GetApplicationConfig() APPConfigContract {

	dataConfig := NewApplicationConfig(map[string]string{
		"port": os.Getenv(envPort),
	})
	return dataConfig
}

// APPConfigContract interface for api data config
type APPConfigContract interface {
	Port() string
}

// NewApplicationConfig return an APPConfigContract implementation
func NewApplicationConfig(envMap map[string]string) APPConfigContract {
	return &appConfig{
		port: envMap["port"],
	}
}

type appConfig struct {
	port string
}

// Port gets port string value
func (ac *appConfig) Port() string {
	return ac.port
}
