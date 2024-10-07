package patterns

import (
	"fmt"
	"patterns/automata"
)

func computePrefixFunction(pattern string) []int {
	m := len(pattern)
	pi := make([]int, m)
	k := 0
	for q := 1; q < m; q++ {
		for k > 0 && pattern[k] != pattern[q] {
			k = pi[k-1]
		}
		if pattern[k] == pattern[q] {
			k++
		}
		pi[q] = k
	}
	return pi
}

// while q > 0 and P[q + 1] ≠ T[i]
// q = π[q]
func automatonTransition(q int, a rune, pattern string, pi []int) int {
	for q > 0 && (q >= len(pattern) || rune(pattern[q]) != a) {
		q = pi[q-1]
	}
	if q < len(pattern) && rune(pattern[q]) == a {
		return q + 1
	}
	return 0
}

func buildKMPAutomaton(pattern string, alphabet string) automata.Automaton {
	m := len(pattern)
	pi := computePrefixFunction(pattern)
	dfa := automata.NewAutomaton()
	for idx := 0; idx <= m; idx++ {
		dfa.AddState(automata.State(fmt.Sprint(idx)))
	}
	dfa.SetStartingState("0")
	dfa.AddAcceptingState(automata.State(fmt.Sprint(m)))

	for idx := 0; idx <= m; idx++ {
		for _, r := range alphabet {
			var nextState int
			if idx < m && rune(pattern[idx]) == r {
				nextState = idx + 1
			} else {
				if idx > 0 {
					nextState = automatonTransition(pi[idx-1], r, pattern, pi)
				} else {
					nextState = 0
				}
			}
			err := dfa.AddTransition(automata.State(fmt.Sprint(idx)), r, automata.State(fmt.Sprint(nextState)))
			if err != nil {
				fmt.Printf("error adding transition: %v", err)
			}
		}
	}
	return dfa
}

// KMPAutomatonMatch searches for the pattern in the text using the KMP automaton
func KMPAutomatonMatch(pattern, text string) (int, error) {
	alphabet := BuildAlphabet(pattern, text)
	dfa := buildKMPAutomaton(pattern, alphabet)

	currentState := automata.State("0")
	count := 0

	// Process the text
	for i, r := range text {
		nextState, err := dfa.GetNextState(currentState, r)
		if err != nil {
			// No valid transition, reset to starting state
			currentState = automata.State("0")
			continue
		}
		currentState = nextState

		if dfa.IsAcceptingState(currentState) {
			// Pattern found at position i - patternLength + 1
			position := i - len(pattern) + 1
			fmt.Printf("Pattern found at position %d\n", position)
			count++
		}
	}
	return count, nil
}
