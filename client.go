package main

import (
	"encoding/json"
	"io"
	"net/http"
)

const apiBaseURL = "https://www.mempool.space/api/"

func fetchData[T any](path string) (*T, error) {
	url := apiBaseURL + path
	resp, err := http.Get(url)
	if err != nil {
		println("Error fetching address data:", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		println("Error: received non-200 response code:", resp.StatusCode)
		return nil, err
	}

	body, _ := io.ReadAll(resp.Body)
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
