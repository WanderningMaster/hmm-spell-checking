package viterbi

import (
	"sort"

	"github.com/WanderningMaster/hmm-spell-checking/internal/hmm"
)

type Path struct {
	Prob   float64
	States []rune
}

func ViterbiKBest(observations []rune, hmm *hmm.HMM, k int) [][]rune {
	n := len(observations)
	if n == 0 {
		return nil
	}

	// Initialize the paths for each state.
	paths := make(map[rune][]Path)
	for state, prob := range hmm.InitProbs {
		if prob > 0 { // Only consider states that can be initial.
			paths[state] = []Path{{Prob: prob * hmm.EmissionProbs[state][observations[0]], States: []rune{state}}}
		}
	}

	// Function to keep top K paths.
	keepTopK := func(paths []Path) []Path {
		if len(paths) > k {
			sort.Slice(paths, func(i, j int) bool {
				return paths[i].Prob > paths[j].Prob // Sort by descending probability.
			})
			return paths[:k]
		}
		return paths
	}

	// Iterate over each observation after the first.
	for i := 1; i < n; i++ {
		newPaths := make(map[rune][]Path)
		for curr, currPaths := range paths {
			for _, path := range currPaths {
				for next, transProb := range hmm.TransitionProbs[curr] {
					emitProb := hmm.EmissionProbs[next][observations[i]]
					newProb := path.Prob * transProb * emitProb
					newPath := make([]rune, len(path.States)+1)
					copy(newPath, path.States)
					newPath[len(newPath)-1] = next
					newPaths[next] = append(newPaths[next], Path{Prob: newProb, States: newPath})
				}
			}
		}

		// Keep only the top K paths for each state.
		for state := range newPaths {
			newPaths[state] = keepTopK(newPaths[state])
		}
		paths = newPaths
	}

	// Collect all paths and sort to find the top K overall.
	var allPaths []Path
	for _, p := range paths {
		allPaths = append(allPaths, p...)
	}
	allPaths = keepTopK(allPaths)

	// Extract the state sequences from the paths.
	result := make([][]rune, len(allPaths))
	for i, path := range allPaths {
		result[i] = path.States
	}
	return result
}
