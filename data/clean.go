package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/WanderningMaster/hmm-spell-checking/internal/logger"
	"github.com/WanderningMaster/hmm-spell-checking/utils"
)

func hasSpecial(str string) bool {
	for _, char := range str {
		if unicode.IsPunct(char) || unicode.IsSymbol(char) || unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func main() {
	logger := logger.GetLogger()

	rHandle, err := os.Open("data/en_keystrokes_pairs.txt")
	wHandle, err := os.Create("data/en_keystrokes_pairs_clean.txt")
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
			words := strings.Fields(line)

			// only substituion and inserrtion
			// not sure how hmm can solve deletion problem
			if len(words[0]) >= len(words[1]) {
				outLen += 1
				wHandle.WriteString(strings.ToLower(line) + "\n")
			}
		}
	}

	logger.Info(
		fmt.Sprintf("Done.\nInitial pairs: %d, Cleaned pairs: %d\n", initialLen, outLen),
	)
}
