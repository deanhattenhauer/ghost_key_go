package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// TODO: bruteForceAttack - generates combinations via itertools equivalent, calls verify
// include schollz/progressbar in bruteforce attack.
// TODO: cuppAttack - collects target profile, generates candidates, calls verify
func dictionaryAttack(storedHash string) (bool, error) {
	color.Yellow("\n[*] Running dictionary attack...\n")
	attemptCounter := 0
	startTime := time.Now()

	// Open rockyou.txt and defer close ensures filed is closed when function returns
	file, err := os.Open("rockyou.txt")
	if err != nil {
		return false, fmt.Errorf("failed to open rockyou: %w", err)
	}
	defer file.Close()

	// Read line by line in the wordlist
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		guess := strings.TrimSpace(line)        // Remove newline character and whitespace
		match, err := verify(guess, storedHash) // Hash guess and compare to storedHash
		if err != nil {
			return false, err
		}
		// If match found, report and exit
		if match {
			color.Green("Password: %s\n", guess)
			fmt.Printf("Attempts: %d\n", attemptCounter)
			fmt.Printf("Time elapsed: %v\n", time.Since(startTime))
			return true, nil
		}
		attemptCounter++
	}
	if err := scanner.Err(); err != nil {
		return false, fmt.Errorf("error reading rockyou: %w", err)
	}
	// Only reaches here if entire wordlist is exhausted with no match
	color.Red("Password not found in wordlist\n")
	fmt.Printf("Attempts: %d\n", attemptCounter)
	fmt.Printf("Time elapsed: %v\n", time.Since(startTime))
	return false, nil
}
