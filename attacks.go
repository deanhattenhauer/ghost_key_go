package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// TODO: dictionaryAttack - reads rockyou.txt line by line, calls verify on each guess
// TODO: bruteForceAttack - generates combinations via itertools equivalent, calls verify
// TODO: cuppAttack - collects target profile, generates candidates, calls verify
func dictionaryAttack(storedHash string) (bool, error) {
	attemptCounter := 0
	startTime := time.Now()

	file, err := os.Open("rockyou.txt")
	if err != nil {
		return false, fmt.Errorf("failed to open rockyou: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		guess := strings.TrimSpace(line)
		match, err := verify(guess, storedHash)
		if err != nil {
			return false, err
		}
		if match {
			fmt.Printf("Password: %s\n", guess)
			fmt.Printf("Attempts: %d\n", attemptCounter)
			fmt.Printf("Time elapsed: %v\n", time.Since(startTime))
			return true, nil
		}
		attemptCounter++
	}
	if err := scanner.Err(); err != nil {
		return false, fmt.Errorf("error reading rockyou: %w", err)
	}
	fmt.Printf("Password not found in wordlist\n")
	fmt.Printf("Attempts: %d\n", attemptCounter)
	fmt.Printf("Time elapsed: %v\n", time.Since(startTime))
	return false, nil
}
