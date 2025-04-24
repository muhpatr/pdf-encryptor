package handler

import (
	"encoding/json"
	"net/http"
	"pdf-encryptor/crypto"
	"pdf-encryptor/logger"
	"time"
)

// ChaChaEncryptHandler godoc
// @Summary Enkripsi file PDF dengan ChaCha20-Poly1305
// @Accept json
// @Produce json
// @Param request body FileRequest true "File info"
// @Success 200 {object} StandardResponse
// @Failure 400 {object} StandardResponse "Bad Request - input tidak valid"
// @Failure 403 {object} StandardResponse "Forbidden - permission denied"
// @Failure 404 {object} StandardResponse "Not Found - file tidak ditemukan"
// @Failure 500 {object} StandardResponse "Internal Server Error"
// @Router /chacha20-poly1305/encrypt [post]
func ChaChaEncryptHandler(w http.ResponseWriter, r *http.Request) {
	var req FileRequest
	json.NewDecoder(r.Body).Decode(&req)

	start := time.Now()
	err := crypto.EncryptChaCha20(req.Source, req.Destination, req.Key)
	duration := time.Since(start)

	logger.LogAction("CHACHA20 ENCRYPT", req.Source, req.Destination, duration, err)
	respond(w, err)
}

// ChaChaDecryptHandler godoc
// @Summary Dekripsi file PDF dengan ChaCha20-Poly1305
// @Accept json
// @Produce json
// @Param request body FileRequest true "File info"
// @Success 200 {object} StandardResponse
// @Failure 400 {object} StandardResponse "Bad Request - input tidak valid"
// @Failure 403 {object} StandardResponse "Forbidden - permission denied"
// @Failure 404 {object} StandardResponse "Not Found - file tidak ditemukan"
// @Failure 500 {object} StandardResponse "Internal Server Error"
// @Router /chacha20-poly1305/decrypt [post]
func ChaChaDecryptHandler(w http.ResponseWriter, r *http.Request) {
	var req FileRequest
	json.NewDecoder(r.Body).Decode(&req)

	start := time.Now()
	err := crypto.DecryptChaCha20(req.Source, req.Destination, req.Key)
	duration := time.Since(start)

	logger.LogAction("CHACHA20 DECRYPT", req.Source, req.Destination, duration, err)
	respond(w, err)
}
