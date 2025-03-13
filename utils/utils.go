package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var Validate = validator.New()

func ParseJSON(req *http.Request, payload any) error {
	if req.Body == nil {
		return fmt.Errorf("missing request body")
	}
	defer req.Body.Close()

	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		return fmt.Errorf("error decoding JSON: %v", err)
	}
	return nil
}

func WriteJSON(rw http.ResponseWriter, status int, payload any) error {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)

	if err := json.NewEncoder(rw).Encode(payload); err != nil {
		return fmt.Errorf("error encoding JSON: %v", err)
	}
	return nil
}

func WriteError(rw http.ResponseWriter, status int, err error) {
	err1 := WriteJSON(rw, status, map[string]string{"error": err.Error()})
	if err1 != nil {
		return
	}
}
