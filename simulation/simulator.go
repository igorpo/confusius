package simulation

import (
	"math/rand"
	"sort"
	"time"

	"github.com/igorpo/confusius/models"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// ScoredOption shows one result of a simulated trial
type ScoredOption struct {
	Name string
	Pct  float64
}

// Simulator represents the simulation world
type Simulator struct {
	choices models.Choices
	counts  map[string]int
	trials  int
	total   int
	results []ScoredOption
}

// New returns a simulation with the desired number of trials
func New(n int, choices models.Choices) *Simulator {
	return &Simulator{
		choices: choices,
		counts:  make(map[string]int),
		trials:  n,
		results: []ScoredOption{},
	}
}

// Simulate runs the simulation for the desired number of trials as given to the simulator
func (s *Simulator) Simulate() {
	for i := 0; i < s.trials; i++ {
		for _, c := range s.choices {
			winner := versus(c.FirstOption, c.SecondOption)
			s.counts[winner]++
		}
	}

	s.total = len(s.counts) * s.trials

	// process results in desc order
	for k, v := range s.counts {
		r := ScoredOption{Name: k, Pct: float64(v) / float64(s.total)}
		s.results = append(s.results, r)
	}

	sort.Slice(s.results, func(i, j int) bool {
		return s.results[i].Pct >= s.results[j].Pct
	})
}

// Top allows for access of the top entries of the simulator. Top must be called after Simulate
// in order for results to be populated.
func (s *Simulator) Top(n int) []ScoredOption {
	if n > len(s.counts) {
		n = len(s.counts)
	}

	return s.results[:n]
}

func versus(opt1 models.Option, opt2 models.Option) string {
	n1 := rand.Intn(opt1.Weight)
	n2 := rand.Intn(opt2.Weight)
	if n1 >= n2 {
		return opt1.Name
	}
	return opt2.Name
}
