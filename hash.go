package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
)

// stringToMD5Hash returns the hex-encoded MD5 hash of password.
// MD5 is cryptographically broken and unsuitable for real security use;
// it's included here only because the original ghost_key supported it
// as a comparison option, not as a recommendation.
func stringToMD5Hash(password string) (string, error) {
	h := md5.New()
	_, err := io.WriteString(h, password)
	if err != nil {
		return "", fmt.Errorf("failed to write password to hasher: %w", err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// stringToSHA256Hash returns the hex-encoded SHA256 hash of a password.
func stringToSHA256Hash(password string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(password))
	if err != nil {
		return "", fmt.Errorf("failed to write password to hasher: %w", err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// TODO: Create verify func to checks guess hasshes to storedHash - auto-detecting
// MD5 vs SHA256 based on the stored hash's length (32 hex chars for MD5, 64 for SHA256).
