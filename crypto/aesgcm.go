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

const encryptionHeader = "PDFCryptMPv1" // ✨ Header untuk validasi file hasil enkripsi

func EncryptAESGCM(source, dest, hexKey string) error {
	key, err := hex.DecodeString(hexKey)
	if err != nil || len(key) != 32 {
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

	// ✨ Tulis header + nonce ke file
	if _, err := output.Write([]byte(encryptionHeader)); err != nil {
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

	// ✨ Cek header dulu
	header := make([]byte, len(encryptionHeader))
	if _, err := io.ReadFull(input, header); err != nil {
		return err
	}
	if string(header) != encryptionHeader {
		return errors.New("file is not encrypted by this system")
	}

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
		// ✨ Return pesan eksplisit agar handler bisa kasih 401
		return errors.New("message authentication failed")
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
