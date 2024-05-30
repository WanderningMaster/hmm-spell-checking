package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/WanderningMaster/hmm-spell-checking/internal/logger"
	"github.com/WanderningMaster/hmm-spell-checking/utils"
)

func hasSpecial(str string) bool {
	for _, char := range str {
		if char == rune('\'') {
			continue
		}
		if unicode.IsPunct(char) || unicode.IsSymbol(char) || unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func cleanVoc() {
	logger := logger.GetLogger()

	rHandle, err := os.Open("data/words.txt")
	wHandle, err := os.Create("data/words_clean.txt")
	defer rHandle.Close()
	defer wHandle.Close()

	utils.Require(err)

	scanner := bufio.NewScanner(rHandle)
	scanner.Split(bufio.ScanLines)

	initialLen := 0
	outLen := 0

	logger.Info("Scanning...")
	for scanner.Scan() {
		line := scanner.Text()
		initialLen += 1

		if !hasSpecial(line) {
			outLen += 1
			wHandle.WriteString(strings.ToLower(line) + "\n")
		}
	}

	logger.Info(
		fmt.Sprintf("Done.\nInitial: %d, Cleaned: %d\n", initialLen, outLen),
	)
}

func splitData(filePath string, trainRatio float64) error {
	// Seed the random number generator to get different results each run
	rand.Seed(time.Now().UnixNano())

	// Open the source file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create training and testing files
	trainFile, err := os.Create("data/training_data.txt")
	if err != nil {
		return fmt.Errorf("error creating training file: %w", err)
	}
	defer trainFile.Close()

	testFile, err := os.Create("data/testing_data.txt")
	if err != nil {
		return fmt.Errorf("error creating testing file: %w", err)
	}
	defer testFile.Close()

	// Use bufio to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Decide randomly where to put the line, according to the trainRatio
		if rand.Float64() < trainRatio {
			_, err = trainFile.WriteString(line + "\n")
			if err != nil {
				return fmt.Errorf("error writing to training file: %w", err)
			}
		} else {
			_, err = testFile.WriteString(line + "\n")
			if err != nil {
				return fmt.Errorf("error writing to testing file: %w", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}

func main() {
	logger := logger.GetLogger()

	rHandle, err := os.Open("data/en_keystrokes_pairs.txt")
	wHandle, err := os.Create("data/en_keystrokes_pairs_clean.txt")
	insertionProblemsHandle, err := os.Create("data/insersition.txt")

	defer rHandle.Close()
	defer wHandle.Close()
	defer insertionProblemsHandle.Close()

	utils.Require(err)

	scanner := bufio.NewScanner(rHandle)
	scanner.Split(bufio.ScanLines)

	initialLen := 0
	outLen := 0

	logger.Info("Scanning...")
	oneDiff := 0
	twoDiff := 0
	moreDiff := 0
	for scanner.Scan() {
		line := scanner.Text()
		initialLen += 1

		if !hasSpecial(line) {
			words := strings.Fields(line)

			if len(words[0]) == len(words[1]) {
				outLen += 1
				wHandle.WriteString(strings.ToLower(line) + "\n")
			}
			if len(words[0]) > len(words[1]) {
				if len(words[1]) < 2 {
					continue
				}
				if len(words[0])-len(words[1]) == 1 {
					oneDiff += 1
				}
				if len(words[0])-len(words[1]) == 2 {
					twoDiff += 1
				}
				if len(words[0])-len(words[1]) > 2 {
					moreDiff += 1
				}
				insertionProblemsHandle.WriteString(strings.ToLower(words[1]) + "\n")
			}
		}
	}
	fmt.Println(oneDiff, twoDiff, moreDiff)

	logger.Info(
		fmt.Sprintf("Done.\nInitial pairs: %d, Cleaned pairs: %d\n", initialLen, outLen),
	)

	cleanVoc()
	// splitData("data/en_keystrokes_pairs_clean.txt", 0.9)
}
