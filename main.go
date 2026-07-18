// Ghost Key recreated in Go.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	result, err := dictionaryAttack(storedHash)
	if err != nil {
		log.Fatal("error with dictionary attack: ", err)
	}
	if result {
		return
	}
	// TODO: fall back to bruteForceAttack here
}
