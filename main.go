// Ghost Key recreated in Go.

package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// bufio.NewReader & ReadString reads everything up to and including the
	// newline character, regardless of spaces (unlike fmt.Scanln, which
	// stops at the first whitespace).
	reader := bufio.NewReader(os.Stdin)

	// Collect target hash from user. Leaving this blank signals "generate
	// a new hash from a password" instead of checking against an existing one.
	fmt.Print("Enter target hash (or press Enter to generate one):")
	storedHash, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("error reading input: ", err)
	}
	storedHash = strings.TrimSpace(storedHash)

	if storedHash == "" {
		// No target hash given, so prompt for a password to hash instead.
		fmt.Print("Enter a password to hash: ")
		password, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("error reading input: ", err)
		}
		password = strings.TrimSpace(password)

		// Don't silently hash an empty string, that's almost certainly
		// a mistake on the user's part, not an intentional empty password.
		if password == "" {
			log.Fatal("no password provided, cannot generate hash")
		}

		fmt.Print("Hash algorithm - MD5 or SHA256 (m / s): ")
		algo, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("error reading input: ", err)
		}

		// Normalize case so "S", "s", " S " etc. all behave the same.
		algo = strings.ToLower(strings.TrimSpace(algo))

		// Defaulkt to MD5 for anything that ins't explicitly "s"; keeps
		// the branching simple, though this does mean typos silently
		// fall back to MD5 rather than re-prompting *** WILL ADD FEATURE FOR REPROMPT ***
		if algo == "s" {
			storedHash, err = stringToSHA256Hash(password)
			if err != nil {
				log.Fatal("error generating hash: ", err)
			}
			fmt.Printf("Generated SHA256 hash: %s\n", storedHash)
		} else {
			storedHash, err = stringToMD5Hash(password)
			if err != nil {
				log.Fatal("error generating hash: ", err)
			}
			fmt.Printf("Generated MD5 hash: %s\n", storedHash)
		}
	}
}

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
