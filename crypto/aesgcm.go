package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

func EncryptAESGCM(source, dest, hexKey string) error {
	key, err := hex.DecodeString(hexKey)
	if err != nil || len(key) != 32 {
		return errors.New("invalid key: must be 256-bit hex string")
	}

	// Buka file sumber
	input, err := os.Open(source)
	if err != nil {
		return err
	}
	defer input.Close()

	// Jika src dan dest sama, gunakan file sementara
	tempPath := dest
	if source == dest {
		tempPath = dest + ".tmp"
	}

	// Buat file tujuan atau sementara
	output, err := os.Create(tempPath)
	if err != nil {
		return err
	}
	defer output.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return err
	}

	if _, err := output.Write(nonce); err != nil {
		return err
	}

	plaintext, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	if _, err := output.Write(ciphertext); err != nil {
		return err
	}

	// Overwrite jika src == dest
	if source == dest {
		output.Close()
		input.Close()
		if err := os.Remove(dest); err != nil {
			return err
		}
		if err := os.Rename(tempPath, dest); err != nil {
			return err
		}
	}

	return nil
}

func DecryptAESGCM(source, dest, hexKey string) error {
	key, err := hex.DecodeString(hexKey)
	if err != nil || len(key) != 32 {
		return errors.New("invalid key: must be 256-bit hex string")
	}

	input, err := os.Open(source)
	if err != nil {
		return err
	}
	defer input.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(input, nonce); err != nil {
		return err
	}

	ciphertext, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	tempPath := dest
	if source == dest {
		tempPath = dest + ".tmp"
	}

	output, err := os.Create(tempPath)
	if err != nil {
		return err
	}
	defer output.Close()

	if _, err := output.Write(plaintext); err != nil {
		return err
	}

	// Overwrite jika src == dest
	if source == dest {
		output.Close()
		input.Close()
		if err := os.Remove(dest); err != nil {
			return err
		}
		if err := os.Rename(tempPath, dest); err != nil {
			return err
		}
	}

	return nil
}
