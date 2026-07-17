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

// verify checks whether guess matches storedHash by hashing the guess using
// the same algorithm as the stored hash. Algorithm is auto-detected by hash
// length: 32 chars = MD5, 64 chars = SHA256.
func verify(guess, storedHash string) (bool, error) {
	var err error
	switch len(storedHash) {
	case 32:
		// MD5 produces a 32 character hex string:
		guess, err = stringToMD5Hash(guess)
		if err != nil {
			return false, fmt.Errorf("failed to hash guess: %w", err)
		}
	case 64:
		// SHA256 produces a 64 character hex string
		guess, err = stringToSHA256Hash(guess)
		if err != nil {
			return false, fmt.Errorf("failed to hash guess: %w", err)
		}
	default:
		// Anything other than 32 or 64 isn't a hash we generated
		return false, fmt.Errorf("unrecognized hash length: %d", len(storedHash))
	}

	// Compare the freshly hased guess aginst the stored hash
	return guess == storedHash, nil
}
