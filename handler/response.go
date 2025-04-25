package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type StandardResponse struct {
	Status  bool   `json:"status" example:"true"`
	Message string `json:"message" example:"Success"`
}

func respond(w http.ResponseWriter, err error) {
	var res StandardResponse
	w.Header().Set("Content-Type", "application/json")

	if err == nil {
		res.Status = true
		res.Message = "Success"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
		return
	}

	// Default response
	res.Status = false
	res.Message = err.Error()

	// Error classification
	switch {
	case strings.Contains(err.Error(), "invalid key"):
		w.WriteHeader(http.StatusBadRequest)

	case strings.Contains(err.Error(), "file is not encrypted by this system"):
		w.WriteHeader(http.StatusUnprocessableEntity) // 422

	case strings.Contains(err.Error(), "message authentication failed"):
		w.WriteHeader(http.StatusBadRequest) // 401

	case os.IsNotExist(err): // file not found
		w.WriteHeader(http.StatusNotFound)

	case os.IsPermission(err): // permission denied
		w.WriteHeader(http.StatusForbidden)

	case isBadRequest(err.Error()):
		w.WriteHeader(http.StatusBadRequest)

	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(res)
}

func isBadRequest(msg string) bool {
	badInputs := []string{
		"invalid", "missing", "unsupported", "not valid", "bad request",
	}
	for _, k := range badInputs {
		if strings.Contains(strings.ToLower(msg), k) {
			return true
		}
	}
	return false
}
