package json

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

// Reads a JSON request into a given data structure.
// This function will return an error if it fails.
func Read(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

// Writes a JSON response to a given http.ResponseWriter
// with the provided status code and data structure.
func Write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Writes a JSON error response to a given http.ResponseWriter
// with the provided status code and message
func WriteError(w http.ResponseWriter, status int, msg string) {
	resp := errorResponse{
		Error: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
