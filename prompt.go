package main

import (
	"bufio"
	"fmt"
	"strings"
)

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
