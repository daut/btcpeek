package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	DefaultApiBaseURL = "https://www.mempool.space/api/"
	DefaultClientType = "api"
)

type Config struct {
	ApiBaseURL string `json:"apiBaseUrl"`
	ClientType string `json:"clientType"`
}

func NewConfig() (*Config, error) {
	config := &Config{
		ApiBaseURL: DefaultApiBaseURL,
		ClientType: DefaultClientType,
	}

	if data, err := os.ReadFile(getConfigPath()); err == nil {
		err = json.Unmarshal(data, config)
		if err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	if apiBaseURL := os.Getenv("BTCPEEK_API_BASE_URL"); apiBaseURL != "" {
		config.ApiBaseURL = apiBaseURL
	}

	if clientType := os.Getenv("BTCPEEK_CLIENT_TYPE"); clientType != "" {
		config.ClientType = clientType
	}

	return config, nil
}

var getConfigPath = func() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(homeDir, ".config", "btcpeek", "config.json")
}
