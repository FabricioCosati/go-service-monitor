package services

import (
	"net/http"
	"time"
)

func HealthCheckService(url string) (int, time.Duration, error) {
	start := time.Now()
	response, err := http.Get(url)

	if err != nil {
		return 0, 0, err
	}

	duration := time.Since(start)
	return response.StatusCode, duration, nil
}
