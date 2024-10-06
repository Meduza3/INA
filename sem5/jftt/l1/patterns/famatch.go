package patterns

import (
	"fmt"
	"patterns/automata"
)

func buildAlphabet(pattern, text string) string {
	charSet := make(map[rune]struct{})
	for _, r := range pattern {
		charSet[r] = struct{}{}
	}
	for _, r := range text {
		charSet[r] = struct{}{}
	}
	var alphabet []rune
	for r := range charSet {
		alphabet = append(alphabet, r)
	}
	return string(alphabet)
}

func AutomatonMatch(pattern, text string) error {
	dfa := automata.NewAutomaton()
	for idx := 0; idx <= len(pattern); idx++ {
		dfa.AddState(automata.State(fmt.Sprint(idx)))
	}
	dfa.SetStartingState("0")
	dfa.AddAcceptingState(automata.State(fmt.Sprint(len(pattern))))

	alphabet := buildAlphabet(pattern, text)

	for idx := 0; idx <= len(pattern); idx++ {
		for _, r := range alphabet {
			s := pattern[:idx] + string(r)
			nextState := automata.State(fmt.Sprint(Suffix(pattern, s)))
			err := dfa.AddTransition(automata.State(fmt.Sprint(idx)), r, nextState)
			if err != nil {
				return fmt.Errorf("error adding transition: %v", err)
			}
		}
	}

	currentState := automata.State("0")

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
		}
	}
	return nil
}
