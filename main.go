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
	storedHash, err := promptInput(reader, "Enter target hash (or press Enter to generate one): ")
	if err != nil {
		log.Fatal("error reading input: ", err)
	}

	if storedHash == "" {
		// No target hash given, so prompt for a password to hash instead.
		password, err := promptInput(reader, "Enter a password to hash: ")
		if err != nil {
			log.Fatal("error reading input: ", err)
		}

		// Don't silently hash an empty string, that's almost certainly
		// a mistake on the user's part, not an intentional empty password.
		if password == "" {
			log.Fatal("no password provided, cannot generate hash")
		}

		// Loops internally until a valid 's' or 'm' is entered.
		algo, err := promptAlgoChoice(reader)
		if err != nil {
			log.Fatal("error reading input: ", err)
		}

		switch algo {
		case "s":
			storedHash, err = stringToSHA256Hash(password)
			if err != nil {
				log.Fatal("error generating hash: ", err)
			}
			fmt.Printf("Generated SHA256 hash: %s\n", storedHash)
		case "m":
			storedHash, err = stringToMD5Hash(password)
			if err != nil {
				log.Fatal("error generating hash: ", err)
			}
			fmt.Printf("Generated MD5 hash: %s\n", storedHash)
		}
	}
}

// promptInput prints prompt, reads a single line from reader, and returns
// it trimmed of surrounding whitespace/newline. Central place for all
// user-input collection so callers don't need to repeat the read/trim/error code.
func promptInput(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading input: %w", err)
	}
	return strings.TrimSpace(input), nil
}

// promptAlgoChoice repeatedly prompts via promptInput until the user enters
// a valid algorithm choice ('s' or 'm'), rather than failing on the first
// invalid input.
func promptAlgoChoice(reader *bufio.Reader) (string, error) {
	for {
		algo, err := promptInput(reader, "Hash algorithm - MD5 or SHA256 (m / s): ")
		if err != nil {
			return "", err
		}
		algo = strings.ToLower(algo)
		if algo == "s" || algo == "m" {
			return algo, nil
		}
		fmt.Println("Invalid choice - please enter 's' or 'm'.")
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
