package statemachine

var (
	NotExistCurrentStateInTransaction      error
	NotExistEventCurrentStateInTransaction error
)

type State int

type Event int

type StateMachine struct {
	initialState State
	currentState State
	transactions map[State]map[Event]State
}

func New(initialState State) *StateMachine {
	return &StateMachine{
		initialState: initialState,
		currentState: initialState,
		transactions: make(map[State]map[Event]State),
	}
}

func (sm *StateMachine) PerformTransition(event Event) error {
	if _, exist := sm.transactions[sm.currentState]; !exist {
		sm.currentState = sm.initialState
		return NotExistCurrentStateInTransaction
	} else if _, exist := sm.transactions[sm.currentState][event]; !exist {
		sm.currentState = sm.initialState
		return NotExistEventCurrentStateInTransaction
	}

	nextState := sm.transactions[sm.currentState][event]
	sm.currentState = nextState

	return nil
}
