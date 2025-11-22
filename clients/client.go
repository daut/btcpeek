package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/daut/btcpeek/config"
)

var httpClient = &http.Client{
	Timeout: 30 * time.Second,
}

type Client struct {
	config *config.Config
}

func NewClient(config *config.Config) *Client {
	return &Client{
		config: config,
	}
}

func (c *Client) FetchData(path string, dest any) error {
	url := c.config.ApiBaseURL + path
	resp, err := httpClient.Get(url)
	if err != nil {
		println("Error fetching address data:", err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return json.NewDecoder(resp.Body).Decode(dest)
}
