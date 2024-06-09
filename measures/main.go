package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/WanderningMaster/hmm-spell-checking/internal/logger"
	"github.com/WanderningMaster/hmm-spell-checking/services"
	"github.com/WanderningMaster/hmm-spell-checking/utils"
)

func getPairs() [][]string {
	rHandle, err := os.Open("data/testing_set.txt")
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

func checkPair(c services.Candidate, actual string) bool {
	if c.Best == actual {
		return true
	}
	for _, c := range c.Variants {
		if c == actual {
			return true
		}
	}

	return false
}

func lambda_accuracy(fileName string, lambda float64) {
	// f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	// utils.Require(err)

	maxConcurrency := 200
	sem := make(chan struct{}, maxConcurrency)
	spellChecker := services.NewSpellChecker(30, lambda)

	testingPairs := getPairs()
	all := len(testingPairs)

	matched := 0
	counter := make(chan struct{}, all)
	wg := sync.WaitGroup{}
	for _, pair := range testingPairs {
		wg.Add(1)
		go spellChecker.CorrectAssert(pair[0], pair[1], counter, sem, &wg)
	}
	wg.Wait()
	close(counter)
	close(sem)

	for range counter {
		matched += 1
	}

	accr := float64(matched) * 100.0 / float64(all)
	fmt.Printf("Lambda %v Matched %d/%d\nAccuracy: %.3f %%\n", lambda, matched, all, accr)
	// f.WriteString(
	// 	fmt.Sprintf("%.3f %.3f\n", accr, lambda),
	// )
}

func variants_accuracy(fileName string, maxVariants int) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	utils.Require(err)
	maxConcurrency := 200
	sem := make(chan struct{}, maxConcurrency)

	lambda := 0.225
	spellChecker := services.NewSpellChecker(maxVariants, lambda)

	testingPairs := getPairs()
	all := len(testingPairs)

	matched := 0
	counter := make(chan struct{}, all)
	wg := sync.WaitGroup{}
	for _, pair := range testingPairs {
		wg.Add(1)
		go spellChecker.CorrectAssert(pair[0], pair[1], counter, sem, &wg)
	}
	wg.Wait()
	close(counter)
	close(sem)

	for range counter {
		matched += 1
	}

	accr := float64(matched) * 100.0 / float64(all)
	fmt.Printf("Max variants %v Matched %d/%d\nAccuracy: %.3f %%\n", maxVariants, matched, all, accr)
	f.WriteString(
		fmt.Sprintf("%.3f %d\n", accr, maxVariants),
	)
}

func measureTime(spellChecker *services.SpellChecker) time.Duration {
	maxConcurrency := 200
	sem := make(chan struct{}, maxConcurrency)

	testingPairs := getPairs()
	all := len(testingPairs)

	matched := 0
	counter := make(chan struct{}, all)
	start := time.Now()
	wg := sync.WaitGroup{}
	for _, pair := range testingPairs {
		wg.Add(1)
		go spellChecker.CorrectAssert(pair[0], pair[1], counter, sem, &wg)
	}
	wg.Wait()
	end := time.Since(start)
	close(counter)
	close(sem)

	for range counter {
		matched += 1
	}

	return end
}
func variants_efficiency(fileName string, maxVariants int) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	utils.Require(err)

	lambda := 0.225
	spellChecker := services.NewSpellChecker(maxVariants, lambda)

	var totalDuration time.Duration
	for range 10 {
		totalDuration += measureTime(spellChecker)
	}
	avrg := totalDuration / time.Duration(10)

	fmt.Printf("Max variants %v \nTime: %v\n", maxVariants, avrg)
	f.WriteString(
		fmt.Sprintf("%d %d\n", avrg.Milliseconds(), maxVariants),
	)
}

func measureTimeSync(spellChecker *services.SpellChecker) time.Duration {
	testingPairs := getPairs()

	matched := 0
	start := time.Now()
	for _, pair := range testingPairs {
		if spellChecker.CorrectAssertSync(pair[0], pair[1]) {
			matched += 1
		}
	}
	end := time.Since(start)

	return end
}
func variants_efficiency_sync(fileName string, maxVariants int) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	utils.Require(err)

	lambda := 0.225
	spellChecker := services.NewSpellChecker(maxVariants, lambda)

	var totalDuration time.Duration
	for range 10 {
		totalDuration += measureTimeSync(spellChecker)
	}
	avrg := totalDuration / time.Duration(10)

	fmt.Printf("Max variants %v \nTime: %v\n", maxVariants, avrg)
	f.WriteString(
		fmt.Sprintf("%d %d\n", avrg.Milliseconds(), maxVariants),
	)
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func main() {
	start := time.Now()
	for v := 43; v <= 50; v++ {
		variants_efficiency_sync("measures/speed_sync.txt", v)
	}
	logger := logger.GetLogger()
	logger.Info(time.Since(start).String())
}
