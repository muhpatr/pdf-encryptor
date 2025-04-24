package handler

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

// KeyResponse digunakan untuk response generate key
type KeyResponse struct {
	Status bool   `json:"status"`
	Key    string `json:"key"`
}

// GenerateKeyHandler godoc
// @Summary Generate 256-bit random encryption key
// @Produce json
// @Success 200 {object} KeyResponse
// @Router /generate-key [get]
func GenerateKeyHandler(w http.ResponseWriter, r *http.Request) {
	key := make([]byte, 32) // 256-bit
	if _, err := rand.Read(key); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "failed to generate key",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(KeyResponse{
		Status: true,
		Key:    hex.EncodeToString(key),
	})
}
