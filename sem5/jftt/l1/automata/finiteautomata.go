package automata

import "fmt"

type Automaton interface {
	AddState(state State)
	SetStartingState(state State) error
	AddAcceptingState(state State) error
	GetNextState(state State, input rune) (State, error)
	AddTransition(from State, input rune, to State) error
	IsAcceptingState(state State) bool
	ProcessInput(input string) bool
}

type State string

type TransitionFunc map[State]map[rune]State

type automaton struct {
	states          map[State]struct{}
	startingState   State
	transitionFunc  TransitionFunc
	acceptingStates map[State]struct{}
}

func NewAutomaton() Automaton {
	return &automaton{
		states:          make(map[State]struct{}),
		transitionFunc:  make(TransitionFunc),
		acceptingStates: make(map[State]struct{}),
	}
}

func (a *automaton) AddState(state State) {
	a.states[state] = struct{}{}
}

func (a *automaton) SetStartingState(state State) error {
	if _, exists := a.states[state]; !exists {
		return fmt.Errorf("state %q does not exist", state)
	}
	a.startingState = state
	return nil
}

func (a *automaton) GetNextState(state State, input rune) (State, error) {
	nextState, ok := a.transitionFunc[state][input]
	if !ok {
		return "", fmt.Errorf("no %v-%v-> transition defined on %v", state, input, a)
	}
	return nextState, nil
}

func (a *automaton) IsAcceptingState(state State) bool {
	for s := range a.acceptingStates {
		if s == state {
			return true
		}
	}
	return false
}

func (a *automaton) AddAcceptingState(state State) error {
	if _, exists := a.states[state]; !exists {
		return fmt.Errorf("state %q does not exist", state)
	}
	a.acceptingStates[state] = struct{}{}
	return nil
}

func (a *automaton) AddTransition(from State, input rune, to State) error {
	if _, exists := a.states[from]; !exists {
		return fmt.Errorf("from state %q does not exist", from)
	}
	if _, exists := a.states[to]; !exists {
		return fmt.Errorf("to state %q does not exist", to)
	}
	if a.transitionFunc[from] == nil {
		a.transitionFunc[from] = make(map[rune]State)
	}
	a.transitionFunc[from][input] = to
	return nil
}

func (a *automaton) ProcessInput(input string) bool {
	state := a.startingState
	for _, r := range input {
		transitions, exists := a.transitionFunc[state]
		if !exists {
			return false // No transitions from current state
		}
		nextState, exists := transitions[r]
		if !exists {
			return false // No transition on input symbol
		}
		state = nextState
	}
	_, isAccepting := a.acceptingStates[state]
	return isAccepting
}
