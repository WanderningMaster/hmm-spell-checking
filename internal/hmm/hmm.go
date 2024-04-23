package hmm

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/WanderningMaster/hmm-spell-checking/utils"
)

type HMM struct {
	data            [][]utils.Tuple
	TransitionProbs map[rune]map[rune]float64
	EmissionProbs   map[rune]map[rune]float64
	InitProbs       map[rune]float64
	cache           bool
}
type HMMConf struct {
	fromCache bool
}

func New(conf HMMConf) *HMM {
	model := &HMM{
		data:            [][]utils.Tuple{},
		TransitionProbs: make(map[rune]map[rune]float64),
		EmissionProbs:   make(map[rune]map[rune]float64),
		InitProbs:       map[rune]float64{},
	}

	return model
}
func WithCache()

func (m *HMM) Load(data []string) {
	for _, s := range data {
		records, _ := utils.MapWordPair(s)
		m.data = append(m.data, records)
	}

	m.calcTransitionMatrix()
	m.calcEmissionMatrix()
	m.calcInitialMatrix()
}

type LogConfig struct {
	Outs       *os.File
	ProbMatrix int
}

func (m *HMM) LogProbs(conf LogConfig) {
	switch conf.ProbMatrix {
	case 1:
		m.LogTransitionMatrix(conf.Outs)
	case 2:
		m.LogEmissionMatrix(conf.Outs)
	case 3:
		m.LogInitializationMatrix(conf.Outs)
	default:
		log.Fatal("unknown prob matrix")
	}
}

func (m *HMM) LogTransitionMatrix(outs *os.File) {
	io.WriteString(outs, "Transition probs:\n")
	for fromState, toStates := range m.TransitionProbs {
		for toState, prob := range toStates {
			io.WriteString(
				outs,
				fmt.Sprintf("%s -> %s - %v\n", strconv.QuoteRune(fromState), strconv.QuoteRune(toState), prob),
			)
		}
	}
}

func (m *HMM) LogEmissionMatrix(outs *os.File) {
	io.WriteString(outs, "Emission probs:\n")
	for fromState, toStates := range m.TransitionProbs {
		for toState, prob := range toStates {
			io.WriteString(
				outs,
				fmt.Sprintf("%s -> %s - %v\n", strconv.QuoteRune(fromState), strconv.QuoteRune(toState), prob),
			)
		}
	}
}

func (m *HMM) LogInitializationMatrix(outs *os.File) {
	io.WriteString(outs, "Initial probs:\n")
	for state, prob := range m.InitProbs {
		io.WriteString(
			outs,
			fmt.Sprintf("%s - %v\n", strconv.QuoteRune(state), prob),
		)
	}
}

func (m *HMM) calcTransitionMatrix() {
	transitionCounts := make(map[rune]map[rune]int)
	for _, seq := range m.data {
		for i := 0; i < len(seq)-1; i += 1 {
			fromState := seq[i].Second
			toState := seq[i+1].Second

			if _, ok := transitionCounts[fromState]; !ok {
				transitionCounts[fromState] = make(map[rune]int)
			}

			if _, ok := transitionCounts[fromState][toState]; !ok {
				transitionCounts[fromState][toState] = 0
			}

			transitionCounts[fromState][toState] += 1
		}
	}

	for fromState, toStates := range transitionCounts {
		totalTransitions := 0
		for _, count := range toStates {
			totalTransitions += count
		}

		m.TransitionProbs[fromState] = make(map[rune]float64)
		for toState, count := range toStates {
			m.TransitionProbs[fromState][toState] = float64(count) / float64(totalTransitions)
		}
	}
}

func (m *HMM) calcEmissionMatrix() {
	emissionCounts := make(map[rune]map[rune]int)
	for _, seq := range m.data {
		for i := 0; i < len(seq); i += 1 {
			observed := seq[i].First
			actual := seq[i].Second

			if _, ok := emissionCounts[actual]; !ok {
				emissionCounts[actual] = make(map[rune]int)
			}

			if _, ok := emissionCounts[actual][observed]; !ok {
				emissionCounts[actual][observed] = 0
			}

			emissionCounts[actual][observed] += 1
		}
	}

	for state, observations := range emissionCounts {
		totalObservations := 0
		for _, count := range observations {
			totalObservations += count
		}

		m.EmissionProbs[state] = make(map[rune]float64)
		for observed, count := range observations {
			m.EmissionProbs[state][observed] = float64(count) / float64(totalObservations)
		}
	}
}

func (m *HMM) calcInitialMatrix() {
	initialCounts := make(map[rune]int)
	for _, seq := range m.data {
		state := seq[0].Second
		if _, ok := initialCounts[state]; !ok {
			initialCounts[state] = 0
		}
		initialCounts[state] += 1
	}

	totalSeqs := len(m.data)
	for state, observations := range initialCounts {
		m.InitProbs[state] = float64(observations) / float64(totalSeqs)
	}
}
