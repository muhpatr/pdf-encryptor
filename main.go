package main

import (
	"net/http"
	_ "pdf-encryptor/docs"
	"pdf-encryptor/handler"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title PDF Encryptor API
// @version 1.0
// @description API untuk enkripsi dan dekripsi file PDF dengan AES-GCM dan ChaCha20-Poly1305
// @host localhost:7082
// @BasePath /
func main() {
	http.HandleFunc("/aes-gcm/encrypt", handler.AesGcmEncryptHandler)
	http.HandleFunc("/aes-gcm/decrypt", handler.AesGcmDecryptHandler)
	http.HandleFunc("/chacha20-poly1305/encrypt", handler.ChaChaEncryptHandler)
	http.HandleFunc("/chacha20-poly1305/decrypt", handler.ChaChaDecryptHandler)
	http.HandleFunc("/generate-key", handler.GenerateKeyHandler)

	http.Handle("/swagger/", httpSwagger.WrapHandler)

	http.ListenAndServe(":7082", nil)
}
