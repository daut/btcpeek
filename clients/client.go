package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiBaseURL = "https://www.mempool.space/api/"

func FetchData[T any](path string) (*T, error) {
	url := apiBaseURL + path
	resp, err := http.Get(url)
	if err != nil {
		println("Error fetching address data:", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result T
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
