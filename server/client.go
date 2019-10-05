package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient struct {
	uri    string
	client *http.Client
}

type HttpError struct {
	StatusCode int
	Message    string
}

func (error *HttpError) Error() string {
	return fmt.Sprintf("HTTP Error. %d - %s", error.StatusCode, error.Message)
}

func NewClient(uri string) *HttpClient {
	c := &http.Client{
		Timeout: time.Second * 60,
	}

	return &HttpClient{
		client: c,
		uri:    uri,
	}
}

type RequestMutateFunc func(r *http.Request) error

func (c *HttpClient) Post(path string, payload, resource interface{}, mutate RequestMutateFunc) (*http.Response, error) {
	return c.request("POST", path, payload, resource, mutate)
}

func (c *HttpClient) Put(path string, payload interface{}, mutate RequestMutateFunc) (*http.Response, error) {
	return c.request("PUT", path, payload, nil, mutate)
}

func (c *HttpClient) Get(path string, resource interface{}, mutate RequestMutateFunc) (*http.Response, error) {
	return c.request("GET", path, nil, resource, mutate)
}

func (c *HttpClient) Delete(path string, mutate RequestMutateFunc) (*http.Response, error) {
	return c.request("DELETE", path, nil, nil, mutate)
}

func (c *HttpClient) Patch(path string, payload interface{}, mutate RequestMutateFunc) (*http.Response, error) {
	return c.request("PATCH", path, payload, nil, mutate)
}

func (c *HttpClient) request(method, path string, payload, resource interface{}, mutate func(r *http.Request) error) (*http.Response, error) {
	endpoint := c.uri + path

	body, err := serialize(method, payload)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, endpoint, body)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/json")

	if mutate != nil {
		if err = mutate(request); err != nil {
			return nil, err
		}
	}

	response, err := c.client.Do(request)

	if err != nil {
		return response, err
	}

	defer response.Body.Close()

	rbody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return response, err
	}

	if response.StatusCode >= 400 {
		errorMessage := string(rbody[:])

		return response, &HttpError{StatusCode: response.StatusCode, Message: errorMessage}
	}

	if resource != nil {
		return response, json.Unmarshal(rbody, resource)
	}

	return response, nil
}

func serialize(method string, payload interface{}) (io.Reader, error) {
	if methodShouldHavePayload(method) && payload != nil {
		serializedPayload, err := json.Marshal(payload)

		if err != nil {
			return nil, err
		}

		return bytes.NewReader(serializedPayload), nil
	}

	return nil, nil
}

func methodShouldHavePayload(method string) bool {
	return method == "PUT" || method == "POST" || method == "PATCH"
}
