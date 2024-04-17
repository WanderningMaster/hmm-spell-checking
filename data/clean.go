package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
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
	rHandle, err := os.Open("en_keystrokes_pairs.txt")
	wHandle, err := os.Create("en_keystrokes_pairs_clean.txt")
	defer rHandle.Close()
	defer wHandle.Close()

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(rHandle)
	scanner.Split(bufio.ScanLines)

	initialLen := 0
	outLen := 0

	fmt.Println("Scanning...")
	for scanner.Scan() {
		line := scanner.Text()
		initialLen += 1
		if !hasSpecial(line) {
			outLen += 1
			wHandle.WriteString(strings.ToLower(line) + "\n")
		}
	}

	fmt.Printf("Done.\nInitial pairs: %d, Cleaned pairs: %d\n", initialLen, outLen)
}
