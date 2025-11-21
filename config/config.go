package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	ApiBaseURL string `json:"apiBaseUrl"`
	ClientType string `json:"clientType"`
}

func NewConfig() *Config {
	config := &Config{
		ApiBaseURL: "https://www.mempool.space/api/",
		ClientType: "api",
	}

	if data, err := os.ReadFile(getConfigPath()); err == nil {
		json.Unmarshal(data, config)
	}

	return config
}

var getConfigPath = func() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(homeDir, ".config", "btcpeek", "config.json")
}
