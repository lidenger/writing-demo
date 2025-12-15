package http_with_generics

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

func HttpGet[T any](url string) (t T, err error) {
	return HttpWithOptions[T](url, "GET", nil, nil, 3*time.Second)
}

func HttpPost[T any](url string, jsonData any) (t T, err error) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	return HttpWithOptions[T](url, "POST", headers, jsonData, 3*time.Second)
}

func HttpWithOptions[T any](url, method string, headers map[string]string, payload any, timeout time.Duration) (t T, err error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &t)
	return
}
