package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func httpClientRequest(request string, url string) ([]byte, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	switch request {
	case "GET":
		req, err := http.NewRequest(request, url, nil)
		if err != nil {
			return nil, fmt.Errorf("error with creating NewRequest: %v", err)
		}

		res, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error with httpClientRequest(client.Do()): %v", err)
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			return nil, fmt.Errorf("error, Status %v", res.Status)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error with ReadAll: %v", err)
		}

		GETCache.Add(url, body)

		return body, nil

	default:
		return nil, fmt.Errorf("invalid request type provided")

	}
}
