package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/WanderningMaster/hmm-spell-checking/services"
	"github.com/WanderningMaster/hmm-spell-checking/utils"
)

func getPairs() [][]string {
	rHandle, err := os.Open("data/testing_data.txt")
	defer rHandle.Close()

	utils.Require(err)

	scanner := bufio.NewScanner(rHandle)
	scanner.Split(bufio.ScanLines)

	pairs := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		pairs = append(pairs, fields)
	}

	return pairs
}

func measure(maxVairiants int, sem chan struct{}, wg *sync.WaitGroup, f io.Writer) {
	defer wg.Done()

	sem <- struct{}{}
	spellChecker := services.NewSpellChecker(maxVairiants)

	testingPairs := getPairs()

	matched := 0
	all := len(testingPairs)
	for _, pair := range testingPairs {
		canidate, _ := spellChecker.Correct(pair[0])
		if canidate.Best == strings.ToLower(pair[1]) {
			matched += 1
		} else {
			for _, c := range canidate.Variants {
				if c == strings.ToLower(pair[1]) {
					matched += 1
				}
			}
		}
	}

	fmt.Println("MAX_VAIRANTS: ", maxVairiants)
	fmt.Fprintf(f, "MAX_VAIRANTS: %d\n", maxVairiants)

	accr := float64(matched) * 100.0 / float64(all)
	fmt.Printf("Matched %d/%d\nAccuracy: %.3f %%\n\n", matched, all, accr)
	fmt.Fprintf(f, "Matched %d/%d\nAccuracy: %.3f %%\n\n", matched, all, accr)

	<-sem
}

func accuracyWithMaxVariants() {
	f, err := os.Create("measures/measures_variants_with_smoothing")
	utils.Require(err)
	maxVariants := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	nWorkers := len(maxVariants)
	maxConcurrency := 2
	sem := make(chan struct{}, maxConcurrency)

	wg := sync.WaitGroup{}
	wg.Add(nWorkers)
	for _, m := range maxVariants {
		go func(m int) {
			measure(m, sem, &wg, f)
		}(m)
	}

	wg.Wait()
	close(sem)
}

func accuracy() {
	f, err := os.Create("measures/measures_smoothing")
	utils.Require(err)
	spellChecker := services.NewSpellChecker(10)

	testingPairs := getPairs()

	matched := 0
	all := len(testingPairs)
	for _, pair := range testingPairs {
		canidate, _ := spellChecker.Correct(pair[0])
		if canidate.Best == strings.ToLower(pair[1]) {
			matched += 1
		}
	}

	accr := float64(matched) * 100.0 / float64(all)
	fmt.Printf("Matched %d/%d\nAccuracy: %.3f %%\n", matched, all, accr)
	fmt.Fprintf(f, "Matched %d/%d\nAccuracy: %.3f %%\n", matched, all, accr)
}

func main() {
	// accuracyWithMaxVariants()
	accuracy()
}
