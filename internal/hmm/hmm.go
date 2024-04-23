package hmm

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/WanderningMaster/hmm-spell-checking/internal/logger"
	"github.com/WanderningMaster/hmm-spell-checking/utils"
)

type HMM struct {
	data            [][]utils.Tuple
	ready           bool
	TransitionProbs map[rune]map[rune]float64
	EmissionProbs   map[rune]map[rune]float64
	InitProbs       map[rune]float64
}

var (
	ModelIsNotReadyYet = errors.New("model is not ready yet")
	CacheNotFound      = errors.New("cache not found")
	FailedToLoadCache  = errors.New("failed to load cache")
)

func New(withFuncs ...func(m *HMM) error) (*HMM, error) {
	model := &HMM{
		ready:           false,
		data:            [][]utils.Tuple{},
		TransitionProbs: make(map[rune]map[rune]float64),
		EmissionProbs:   make(map[rune]map[rune]float64),
		InitProbs:       map[rune]float64{},
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

func (m *HMM) Load(data []string) {
	for _, s := range data {
		records, _ := utils.MapWordPair(s)
		m.data = append(m.data, records)
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
		return fmt.Errorf("unknown prob matrix")
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
			m.TransitionProbs[fromState][toState] = float64(count) / float64(totalTransitions)
		}
	}
}

func (m *HMM) calcEmissionMatrix() {
	emissionCounts := make(map[rune]map[rune]int)
	for _, seq := range m.data {
		for i := 0; i < len(seq); i += 1 {
			observed := seq[i].Observed
			state := seq[i].State

			if _, ok := emissionCounts[state]; !ok {
				emissionCounts[state] = make(map[rune]int)
			}

			if _, ok := emissionCounts[state][observed]; !ok {
				emissionCounts[state][observed] = 0
			}

			emissionCounts[state][observed] += 1
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
