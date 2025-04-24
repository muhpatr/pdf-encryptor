package handler

import (
	"encoding/json"
	"net/http"
	"pdf-encryptor/crypto"
	"pdf-encryptor/logger"
	"time"
)

type FileRequest struct {
	Source      string `json:"src"`
	Destination string `json:"dest"`
	Key         string `json:"key"`
}

// AesGcmEncryptHandler godoc
// @Summary Enkripsi file PDF dengan AES-GCM
// @Accept json
// @Produce json
// @Param request body FileRequest true "File info"
// @Success 200 {object} StandardResponse
// @Failure 400 {object} StandardResponse "Bad Request - input tidak valid"
// @Failure 403 {object} StandardResponse "Forbidden - permission denied"
// @Failure 404 {object} StandardResponse "Not Found - file tidak ditemukan"
// @Failure 500 {object} StandardResponse "Internal Server Error"
// @Router /aes-gcm/encrypt [post]
func AesGcmEncryptHandler(w http.ResponseWriter, r *http.Request) {
	var req FileRequest
	json.NewDecoder(r.Body).Decode(&req)

	start := time.Now()
	err := crypto.EncryptAESGCM(req.Source, req.Destination, req.Key)

	duration := time.Since(start)

	logger.LogAction("AES-GCM ENCRYPT", req.Source, req.Destination, duration, err)
	respond(w, err)
}

// AesGcmDecryptHandler godoc
// @Summary Dekripsi file PDF dengan AES-GCM
// @Accept json
// @Produce json
// @Param request body FileRequest true "File info"
// @Success 200 {object} StandardResponse
// @Failure 400 {object} StandardResponse "Bad Request - input tidak valid"
// @Failure 403 {object} StandardResponse "Forbidden - permission denied"
// @Failure 404 {object} StandardResponse "Not Found - file tidak ditemukan"
// @Failure 500 {object} StandardResponse "Internal Server Error"
// @Router /aes-gcm/decrypt [post]
func AesGcmDecryptHandler(w http.ResponseWriter, r *http.Request) {
	var req FileRequest
	json.NewDecoder(r.Body).Decode(&req)

	start := time.Now()
	err := crypto.DecryptAESGCM(req.Source, req.Destination, req.Key)
	duration := time.Since(start)

	logger.LogAction("AES-GCM DECRYPT", req.Source, req.Destination, duration, err)
	respond(w, err)
}
