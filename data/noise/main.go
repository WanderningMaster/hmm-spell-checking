package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/WanderningMaster/hmm-spell-checking/utils"
)

var keyboardNeighbors = map[rune]string{
	'a': "aqwsz",
	'b': "bvghn",
	'c': "cxdfv",
	'd': "dserfcx",
	'e': "ewsdr",
	'f': "fdertgvc",
	'g': "gftyhbv",
	'h': "hgyujnb",
	'i': "iujko",
	'j': "jhuikmn",
	'k': "kjiolm",
	'l': "lkop",
	'm': "mnjk",
	'n': "nbhjm",
	'o': "oiklp",
	'p': "pol",
	'q': "qwa",
	'r': "redft",
	's': "swedxza",
	't': "trfgy",
	'u': "uyhji",
	'v': "vcfgb",
	'w': "wqase",
	'x': "xzsdc",
	'y': "ytghu",
	'z': "zasx",
	' ': " ", // Including space for completeness
}

func noisyWord(word string, insertProbability float64, maxInsertions int) (string, string) {
	rand.Seed(time.Now().UnixNano())
	whichNoisy := []bool{}
	result := ""
	insertions := 0
	for _, char := range word {
		result += string(char)
		if insertions < maxInsertions && rand.Float64() < insertProbability {
			if neighbors, ok := keyboardNeighbors[char]; ok {
				// Pick a random neighbor
				index := rand.Intn(len(neighbors))
				result += string(neighbors[index])
				insertions += 1
			}
			whichNoisy = append(whichNoisy, false)
			whichNoisy = append(whichNoisy, true)
		} else {
			whichNoisy = append(whichNoisy, false)
		}
	}

	normalizedWord := ""
	idx := 0
	for _, noisy := range whichNoisy {
		if noisy {
			normalizedWord += "#"
		} else {
			normalizedWord += string(word[idx])
			idx += 1
		}
	}
	return result, normalizedWord
}

func main() {
	rHandle, err := os.Open("data/insersition.txt")
	wHandle, err := os.Create("data/insersition_prep.txt")
	defer rHandle.Close()
	defer wHandle.Close()

	utils.Require(err)

	scanner := bufio.NewScanner(rHandle)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		noisy := ""
		normalized := ""
		for {
			noisy, normalized = noisyWord(line, 0.4, 2)
			if noisy != line {
				break
			}
		}
		wHandle.WriteString(fmt.Sprintf("%s %s\n", noisy, normalized))
	}
}
