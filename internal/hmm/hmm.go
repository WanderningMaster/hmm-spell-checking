package hmm

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"

	"github.com/WanderningMaster/hmm-spell-checking/internal/logger"
	"github.com/WanderningMaster/hmm-spell-checking/utils"
)

type Coord struct {
	X, Y float64
}

type HMM struct {
	data            [][]utils.Tuple
	insertionData   [][]utils.Tuple
	ready           bool
	TransitionProbs map[rune]map[rune]float64
	EmissionProbs   map[rune]map[rune]float64
	InitProbs       map[rune]float64
	KeyboardLayout  map[rune]Coord
}

var (
	ModelIsNotReadyYet = errors.New("model is not ready yet")
	CacheNotFound      = errors.New("cache not found")
	FailedToLoadCache  = errors.New("failed to load cache")
	UnknownProbMatrix  = errors.New("unknown prob matrix")
)

func New(withFuncs ...func(m *HMM) error) (*HMM, error) {
	model := &HMM{
		ready:           false,
		data:            [][]utils.Tuple{},
		insertionData:   [][]utils.Tuple{},
		TransitionProbs: make(map[rune]map[rune]float64),
		EmissionProbs:   make(map[rune]map[rune]float64),
		InitProbs:       map[rune]float64{},
		KeyboardLayout:  InitKeyboardLayout(),
	}
	for _, fn := range withFuncs {
		err := fn(model)
		if err != nil {
			return nil, err
		}
	}

	return model, nil
}
func WithCache(m *HMM) error {
	logger := logger.GetLogger()
	logger.Info("Loading model from cache...")

	buff, err := os.ReadFile("cache/model.bin")
	if err != nil {
		return CacheNotFound
	}
	var b bytes.Buffer
	_, err = b.Write(buff)
	utils.Require(err)

	dec := gob.NewDecoder(&b)
	err = dec.Decode(m)

	if err != nil {
		return fmt.Errorf("%s: %v", FailedToLoadCache.Error(), err)
	}
	m.ready = true

	return nil
}

func (k *HMM) KeyDistance(a, b rune) float64 {
	coordA, okA := k.KeyboardLayout[a]
	coordB, okB := k.KeyboardLayout[b]
	if !okA || !okB {
		return math.Inf(1)
	}
	return math.Sqrt(math.Pow(coordA.X-coordB.X, 2) + math.Pow(coordA.Y-coordB.Y, 2))
}

func InitKeyboardLayout() map[rune]Coord {
	// Initialize the KeyboardLayout map
	return map[rune]Coord{
		// Top row
		'q': {X: 0, Y: 1}, 'w': {X: 1, Y: 1}, 'e': {X: 2, Y: 1}, 'r': {X: 3, Y: 1}, 't': {X: 4, Y: 1},
		'y': {X: 5, Y: 1}, 'u': {X: 6, Y: 1}, 'i': {X: 7, Y: 1}, 'o': {X: 8, Y: 1}, 'p': {X: 9, Y: 1},

		// Home row
		'a': {X: 0, Y: 2}, 's': {X: 1, Y: 2}, 'd': {X: 2, Y: 2}, 'f': {X: 3, Y: 2}, 'g': {X: 4, Y: 2},
		'h': {X: 5, Y: 2}, 'j': {X: 6, Y: 2}, 'k': {X: 7, Y: 2}, 'l': {X: 8, Y: 2}, '\'': {X: 10, Y: 2},

		// Bottom row
		'z': {X: 0, Y: 3}, 'x': {X: 1, Y: 3}, 'c': {X: 2, Y: 3}, 'v': {X: 3, Y: 3}, 'b': {X: 4, Y: 3},
		'n': {X: 5, Y: 3}, 'm': {X: 6, Y: 3},
	}
}

func (m *HMM) Load(data []string, insertionData []string) {
	for _, s := range data {
		records, _ := utils.MapWordPair(s)
		m.data = append(m.data, records)
	}
	for _, s := range insertionData {
		records, _ := utils.MapWordPair(s)
		m.insertionData = append(m.insertionData, records)
	}

	m.calcTransitionMatrix()
	m.calcEmissionMatrix()
	m.calcInitialMatrix()
	m.CacheModel()

	m.ready = true
}

func (m *HMM) CacheModel() {
	file, err := os.Create("cache/model.bin")
	utils.Require(err)
	defer file.Close()

	var b bytes.Buffer

	enc := gob.NewEncoder(&b)
	err = enc.Encode(*m)
	utils.Require(err)

	file.Write(b.Bytes())
}

type LogConfig struct {
	Outs       *os.File
	ProbMatrix int
}

func (m *HMM) LogProbs(conf LogConfig) error {
	if !m.ready {
		return ModelIsNotReadyYet
	}
	switch conf.ProbMatrix {
	case 1:
		m.LogTransitionMatrix(conf.Outs)
	case 2:
		m.LogEmissionMatrix(conf.Outs)
	case 3:
		m.LogInitializationMatrix(conf.Outs)
	default:
		return UnknownProbMatrix
	}

	return nil
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
	for fromState, toStates := range m.EmissionProbs {
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
			fromState := seq[i].State
			toState := seq[i+1].State

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
			m.TransitionProbs[fromState][toState] = (float64(count) / float64(totalTransitions))
		}
	}
}

func (m *HMM) calcEmissionMatrix() {
	rawEmissionCounts := make(map[rune]map[rune]int)
	for _, seq := range m.data {
		for i := 0; i < len(seq); i += 1 {
			observed := seq[i].Observed
			state := seq[i].State

			if _, ok := rawEmissionCounts[state]; !ok {
				rawEmissionCounts[state] = make(map[rune]int)
			}

			if _, ok := rawEmissionCounts[state][observed]; !ok {
				rawEmissionCounts[state][observed] = 0
			}

			rawEmissionCounts[state][observed] += 1
		}
	}

	m.EmissionProbs = make(map[rune]map[rune]float64)
	lambda := 0.035
	for state, observations := range rawEmissionCounts {
		totalObservations := 0
		for _, count := range observations {
			totalObservations += count
		}

		m.EmissionProbs[state] = make(map[rune]float64)
		for observed, count := range observations {
			countProb := float64(count) / float64(totalObservations)

			distance := m.KeyDistance(state, observed)
			distanceProb := math.Exp(-lambda * distance)

			m.EmissionProbs[state][observed] = countProb * distanceProb
			// m.EmissionProbs[state][observed] = countProb
		}
	}
}

func (m *HMM) calcInitialMatrix() {
	initialCounts := make(map[rune]int)
	for _, seq := range m.data {
		state := seq[0].State
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
