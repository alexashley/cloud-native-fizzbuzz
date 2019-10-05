package main_test

import (
	"github.com/alexashley/cloud-native-fizzbuzz/domain"
	"github.com/alexashley/cloud-native-fizzbuzz/server"
	"net/http"
	"testing"
)

func TestConversion(t *testing.T) {
	c := server.NewClient("http://localhost:8080")

	payload := domain.ItoaRequest{Integer: 123}
	var result domain.ItoaResponse

	response, err := c.Post("/api/v1/str/itoa", payload, &result, nil)
	if err != nil {
		t.Error(err)
	}

	if result.String != "123" {
		t.Errorf("Expected 123, got %s", result.String)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", response.StatusCode)
	}
}


func TestEmptyConversion(t *testing.T) {
	c := server.NewClient("http://localhost:8080")

	payload := domain.ItoaRequest{}
	var result domain.ItoaResponse

	_, err := c.Post("/api/v1/str/itoa", payload, &result, nil)
	if err != nil {
		t.Error(err)
	}

	if result.String != "0" {
		t.Errorf("Expected 0, got %s", result.String)
	}
}

func TestInvalidPayload(t *testing.T) {
	c := server.NewClient("http://localhost:8080")

	payload := struct {
		Integer string
	}{Integer: "123"}
	var result domain.ItoaResponse

	_, err := c.Post("/api/v1/str/itoa", payload, &result, nil)

	if err == nil {
		t.Error("Expected http error")
	}

	if err, ok := err.(*server.HttpError); ok {
		if err.StatusCode != http.StatusBadRequest {
			t.Error("Expected status to be bad request")
		}
	} else {
		t.Error(err)
	}
}