package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiBaseURL = "https://www.mempool.space/api/"

func FetchData(path string, dest any) error {
	url := apiBaseURL + path
	resp, err := http.Get(url)
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
