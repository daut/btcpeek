package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name          string
		configContent map[string]string
		expected      *Config
	}{
		{
			name:          "returns default config when no config file exists",
			configContent: nil,
			expected: &Config{
				ApiBaseURL: "https://www.mempool.space/api/",
				ClientType: "api",
			},
		},
		{
			name: "partially overrides defaults",
			configContent: map[string]string{
				"apiBaseUrl": "https://www.custom.api/",
			},
			expected: &Config{
				ApiBaseURL: "https://www.custom.api/",
				ClientType: "api",
			},
		},
		{
			name: "fully overrides defaults",
			configContent: map[string]string{
				"apiBaseUrl": "https://localhost:8332",
				"clientType": "rpc",
			},
			expected: &Config{
				ApiBaseURL: "https://localhost:8332",
				ClientType: "rpc",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var configPath string
			if tt.configContent != nil {
				tempDir := t.TempDir()
				configPath = filepath.Join(tempDir, "config.json")
				data, _ := json.Marshal(tt.configContent)
				os.WriteFile(configPath, data, 0644)

				originalGetConfigPath := getConfigPath
				getConfigPath = func() string { return configPath }
				t.Cleanup(func() { getConfigPath = originalGetConfigPath })
			}

			config := NewConfig()

			if config.ApiBaseURL != tt.expected.ApiBaseURL {
				t.Errorf("expected ApiBaseURL %s, got %s", tt.expected.ApiBaseURL, config.ApiBaseURL)
			}
			if config.ClientType != tt.expected.ClientType {
				t.Errorf("expected ClientType %s, got %s", tt.expected.ClientType, config.ClientType)
			}
		})
	}
}
