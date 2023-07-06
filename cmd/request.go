package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

func makeRequest(_ context.Context, uri string, method string, payload interface{}) (interface{}, error) {
	body, jerr := json.Marshal(payload)

	if jerr != nil {
		return nil, jerr
	}

	req, err := http.NewRequest(method, baseURL+uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": os.Getenv("BASIC_AUTH"),
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	responseBody := map[string]interface{}{}

	unjErr := json.Unmarshal(respBody, &responseBody)

	if unjErr != nil {
		return nil, unjErr
	}

	return responseBody, nil
}
