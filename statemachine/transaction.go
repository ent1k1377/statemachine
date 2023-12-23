package statemachine

type Transaction struct {
	CurrentState State
	Transitions  []Transition
}

type Transition struct {
	Event       Event
	Transaction *Transaction
}

func (sm *StateMachine) SetNestedTransaction(transaction *Transaction) {
	if transaction.Transitions == nil {
		return
	}
	for _, transition := range transaction.Transitions {
		q := make(map[Event]State)
		q[transition.Event] = transition.Transaction.CurrentState
		sm.setTransaction(transaction.CurrentState, q)
	}
}

func (sm *StateMachine) SetNonNestedTransaction(state State, transition map[Event]State) {
	sm.setTransaction(state, transition)
}

func (sm *StateMachine) setTransaction(state State, nextStates map[Event]State) {
	sm.transactions[state] = nextStates
}
