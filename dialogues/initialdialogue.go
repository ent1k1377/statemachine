package dialogues

import (
	"github.com/ent1k1377/statemachine/statemachine"
)

const (
	InitialState statemachine.State = iota + 100
	SendBotDescription
)

func NewInitialDialogue(q *statemachine.StateMachine) {
	q.SetNonNestedTransaction(InitialState, map[statemachine.Event]statemachine.State{
		StartCommand:      SendBotDescription,
		HelpCommand:       SendBotDescription,
		NewServiceCommand: WaitingForServiceName,
	})
}
