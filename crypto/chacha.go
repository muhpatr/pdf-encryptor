package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"os"

	"golang.org/x/crypto/chacha20poly1305"
)

func EncryptChaCha20(source, dest, hexKey string) error {
	key, err := hex.DecodeString(hexKey)
	if err != nil || len(key) != chacha20poly1305.KeySize {
		return errors.New("invalid key: must be 256-bit hex string")
	}

	input, err := os.Open(source)
	if err != nil {
		return err
	}
	defer input.Close()

	tempPath := dest
	if source == dest {
		tempPath = dest + ".tmp"
	}

	output, err := os.Create(tempPath)
	if err != nil {
		return err
	}
	defer output.Close()

	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return err
	}

	nonce := make([]byte, chacha20poly1305.NonceSize)
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

	ciphertext := aead.Seal(nil, nonce, plaintext, nil)
	if _, err := output.Write(ciphertext); err != nil {
		return err
	}

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

func DecryptChaCha20(source, dest, hexKey string) error {
	key, err := hex.DecodeString(hexKey)
	if err != nil || len(key) != chacha20poly1305.KeySize {
		return errors.New("invalid key: must be 256-bit hex string")
	}

	input, err := os.Open(source)
	if err != nil {
		return err
	}
	defer input.Close()

	nonce := make([]byte, chacha20poly1305.NonceSize)
	if _, err := io.ReadFull(input, nonce); err != nil {
		return err
	}

	ciphertext, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return err
	}

	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
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
