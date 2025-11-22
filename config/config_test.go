package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name     string
		setup    func()
		expected *Config
	}{
		{
			name:  "returns default config when no config file exists",
			setup: func() {},
			expected: &Config{
				ApiBaseURL: DefaultApiBaseURL,
				ClientType: DefaultClientType,
			},
		},
		{
			name: "partially overrides defaults",
			setup: func() {
				configPath := createFakeConfFile(t, map[string]string{
					"apiBaseUrl": "https://www.custom.api/",
				})
				originalGetConfigPath := getConfigPath
				getConfigPath = func() string { return configPath }
				t.Cleanup(func() { getConfigPath = originalGetConfigPath })
			},
			expected: &Config{
				ApiBaseURL: "https://www.custom.api/",
				ClientType: DefaultClientType,
			},
		},
		{
			name: "fully overrides defaults",
			setup: func() {
				configPath := createFakeConfFile(t, map[string]string{
					"apiBaseUrl": "https://localhost:8332",
					"clientType": "rpc",
				})
				originalGetConfigPath := getConfigPath
				getConfigPath = func() string { return configPath }
				t.Cleanup(func() { getConfigPath = originalGetConfigPath })
			},
			expected: &Config{
				ApiBaseURL: "https://localhost:8332",
				ClientType: "rpc",
			},
		},
		{
			name: "env var overrides config",
			setup: func() {
				configPath := createFakeConfFile(t, nil)
				t.Setenv("BTCPEEK_API_BASE_URL", "https://env.localhost:8332")
				t.Setenv("BTCPEEK_CLIENT_TYPE", "envrpc")
				originalGetConfigPath := getConfigPath
				getConfigPath = func() string { return configPath }
				t.Cleanup(func() { getConfigPath = originalGetConfigPath })
			},
			expected: &Config{
				ApiBaseURL: "https://env.localhost:8332",
				ClientType: "envrpc",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			config, err := NewConfig()

			if err != nil {
				t.Errorf("expected nil, got %s error", err)
			}
			if config.ApiBaseURL != tt.expected.ApiBaseURL {
				t.Errorf("expected ApiBaseURL %s, got %s", tt.expected.ApiBaseURL, config.ApiBaseURL)
			}
			if config.ClientType != tt.expected.ClientType {
				t.Errorf("expected ClientType %s, got %s", tt.expected.ClientType, config.ClientType)
			}
		})
	}
}

func createFakeConfFile(t *testing.T, content map[string]string) string {
	if content == nil {
		content = make(map[string]string)
	}
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.json")
	data, _ := json.Marshal(content)
	os.WriteFile(configPath, data, 0644)
	return configPath
}
