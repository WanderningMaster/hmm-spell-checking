package viterbi

import (
	"math"

	"github.com/WanderningMaster/hmm-spell-checking/internal/hmm"
)

func Viterbi(observations []rune, hmm *hmm.HMM) []rune {
	// Number of observations
	n := len(observations)
	// States of the HMM
	states := make([]rune, 0, len(hmm.InitProbs))
	for k := range hmm.InitProbs {
		states = append(states, k)
	}

	// DP tables for probabilities and back-pointers
	viterbi := make([]map[rune]float64, n)
	backpointer := make([]map[rune]rune, n)

	// Initialization step
	viterbi[0] = make(map[rune]float64)
	backpointer[0] = make(map[rune]rune)
	for _, s := range states {
		viterbi[0][s] = math.Log(hmm.InitProbs[s]) + math.Log(hmm.EmissionProbs[s][observations[0]])
		backpointer[0][s] = 0 // No previous state
	}

	// Recursion step
	for t := 1; t < n; t++ {
		viterbi[t] = make(map[rune]float64)
		backpointer[t] = make(map[rune]rune)
		for _, s := range states {
			maxProb := math.Inf(-1)
			var maxState rune
			for _, sp := range states {
				prob := viterbi[t-1][sp] + math.Log(hmm.TransitionProbs[sp][s]) + math.Log(hmm.EmissionProbs[s][observations[t]])
				if prob > maxProb {
					maxProb = prob
					maxState = sp
				}
			}
			viterbi[t][s] = maxProb
			backpointer[t][s] = maxState
		}
	}

	// Termination step to find the last state
	lastState := states[0]
	maxProb := viterbi[n-1][lastState]
	for _, s := range states {
		if viterbi[n-1][s] > maxProb {
			maxProb = viterbi[n-1][s]
			lastState = s
		}
	}

	// Path backtracking
	path := make([]rune, n)
	path[n-1] = lastState
	for t := n - 1; t > 0; t-- {
		path[t-1] = backpointer[t][path[t]]
	}

	return path
}
