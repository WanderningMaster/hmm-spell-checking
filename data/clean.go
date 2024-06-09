package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
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
	// Open the dataset file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	// Read the lines
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	// Shuffle the lines
	rand.Shuffle(len(lines), func(i, j int) { lines[i], lines[j] = lines[j], lines[i] })

	// Split into training and testing sets (80% training, 20% testing)
	splitIndex := int(trainRatio * float64(len(lines)))
	trainingSet := lines[:splitIndex]
	testingSet := lines[splitIndex:]

	// Save the training set
	trainingFile, err := os.Create("data/training_set.txt")
	if err != nil {
		fmt.Println("Error creating training set file:", err)
		return err
	}
	defer trainingFile.Close()
	for _, line := range trainingSet {
		trainingFile.WriteString(line + "\n")
	}

	// Save the testing set
	testingFile, err := os.Create("data/testing_set.txt")
	if err != nil {
		fmt.Println("Error creating testing set file:", err)
		return err
	}
	defer testingFile.Close()
	for _, line := range testingSet {
		testingFile.WriteString(line + "\n")
	}

	fmt.Println("Dataset split completed successfully.")

	return nil
}

// func main() {
// 	logger := logger.GetLogger()
//
// 	rHandle, err := os.Open("data/en_keystrokes_pairs.txt")
// 	wHandle, err := os.Create("data/en_keystrokes_pairs_clean.txt")
// 	insertionProblemsHandle, err := os.Create("data/insersition.txt")
//
// 	defer rHandle.Close()
// 	defer wHandle.Close()
// 	defer insertionProblemsHandle.Close()
//
// 	utils.Require(err)
//
// 	scanner := bufio.NewScanner(rHandle)
// 	scanner.Split(bufio.ScanLines)
//
// 	initialLen := 0
// 	outLen := 0
//
// 	logger.Info("Scanning...")
// 	oneDiff := 0
// 	twoDiff := 0
// 	moreDiff := 0
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		initialLen += 1
//
// 		if !hasSpecial(line) {
// 			words := strings.Fields(line)
//
// 			if len(words[0]) == len(words[1]) {
// 				outLen += 1
// 				wHandle.WriteString(strings.ToLower(line) + "\n")
// 			}
// 			if len(words[0]) > len(words[1]) {
// 				if len(words[1]) < 2 {
// 					continue
// 				}
// 				if len(words[0])-len(words[1]) == 1 {
// 					oneDiff += 1
// 				}
// 				if len(words[0])-len(words[1]) == 2 {
// 					twoDiff += 1
// 				}
// 				if len(words[0])-len(words[1]) > 2 {
// 					moreDiff += 1
// 				}
// 				insertionProblemsHandle.WriteString(strings.ToLower(words[1]) + "\n")
// 			}
// 		}
// 	}
// 	fmt.Println(oneDiff, twoDiff, moreDiff)
//
// 	logger.Info(
// 		fmt.Sprintf("Done.\nInitial pairs: %d, Cleaned pairs: %d\n", initialLen, outLen),
// 	)
//
// 	cleanVoc()
// 	// splitData("data/en_keystrokes_pairs_clean.txt", 0.9)
// }

func main() {
	splitData("data/en_keystrokes_pairs_clean.txt", 0.8)
}
