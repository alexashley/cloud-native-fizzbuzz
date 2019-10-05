package main_test

import (
	"github.com/alexashley/cloud-native-fizzbuzz/domain"
	"github.com/alexashley/cloud-native-fizzbuzz/server"
	"net/http"
	"testing"
)

func TestDivisibleByThree(t *testing.T) {
	c := server.NewClient("http://localhost:8080")

	payload := domain.Mod3Query{Value: 3}
	var result domain.Mod3Result

	response, err := c.Post("/api/v1/math/mod/3", payload, &result, nil)
	if err != nil {
		t.Error(err)
	}

	if !result.IsDivisibleByThree {
		t.Error("Should be divisible by three")
	}

	if response.StatusCode != 200 {
		t.Error("Status should be OK")
	}
}

func TestNotDivisbleByThree(t *testing.T) {
	c := server.NewClient("http://localhost:8080")

	payload := domain.Mod3Query{Value: 4}
	var result domain.Mod3Result

	response, err := c.Post("/api/v1/math/mod/3", payload, &result, nil)

	if err != nil {
		t.Error(err)
	}

	if result.IsDivisibleByThree {
		t.Error("Should **not** be divisible by three")
	}

	if response.StatusCode != http.StatusOK {
		t.Error("Status should be OK")
	}
}

func TestInvalidPayload(t *testing.T) {
	c := server.NewClient("http://localhost:8080")
	payload := struct{ Value string }{
		"abc",
	}
	var result domain.Mod3Result

	_, err := c.Post("/api/v1/math/mod/3", payload, &result, nil)

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

func TestEmptyFieldResult(t *testing.T) {
	c := server.NewClient("http://localhost:8080")
	payload := struct{ Foobar string }{ // No value field, so will be evaluated as 0
		"abc",
	}

	var result domain.Mod3Result

	_, err := c.Post("/api/v1/math/mod/3", payload, &result, nil)

	if err != nil {
		t.Error(err)
	}

	if !result.IsDivisibleByThree {
		t.Error("Should be divisible by three")
	}
}
