package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestSetOperation(t *testing.T) {
	payload := map[string]string{
		"key":   "test-key",
		"value": "test-value",
	}
	payloadBytes, _ := json.Marshal(payload)

	resp, err := http.Post("http://localhost:8000/set", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatalf("Failed to make POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

}

func TestGetOperation(t *testing.T) {
	key := "abc-1"

	resp, err := http.Get("http://localhost:8000/get/" + key)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

}

func TestSearchOperation(t *testing.T) {
	queryParam := "abc"

	resp, err := http.Get("http://localhost:8000/search?prefix=" + queryParam)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

}
