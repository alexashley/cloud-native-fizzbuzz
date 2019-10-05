package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		BadRequest(w)
		Warnf(r.Context(), "Failed to read body: %v", err)

		return false
	}

	Debugf(r.Context(), "body: %s", body)

	if err = json.Unmarshal(body, v); err != nil {
		Warnf(r.Context(), "Failed to parse body: %v", err)
		BadRequest(w)
		return false
	}

	return true
}

func WriteBody(w http.ResponseWriter, r *http.Request, v interface{}) {
	res, err := json.Marshal(v)

	if err != nil {
		Internal(w)
		Warnf(r.Context(), "Failed to serialize body: %v", err)
		return
	}

	if written, err := w.Write(res); err != nil || written != len(res) {
		Internal(w)
		Warnf(r.Context(), "Error writing response %v", err)
		return
	}
}
