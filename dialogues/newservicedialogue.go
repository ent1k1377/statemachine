package dialogues

import (
	"github.com/ent1k1377/statemachine/statemachine"
)

// Константы Event
const (
	StartCommand statemachine.Event = iota
	HelpCommand
	CustomMessage
	YesCommand
	NoCommand
	UnknownCommand
	NewServiceCommand
)

const (
	WaitingForServiceName statemachine.State = iota
	WaitingForPasswordExist
	WaitingForPassword
	ChoiceCharactersNumber
	ChoiceCharactersAlphabet
)

func NewServiceDialogue() statemachine.Transaction {
	return statemachine.Transaction{
		CurrentState: WaitingForServiceName,
		Transitions:  waitingForServiceNameTransition(),
	}
}

func waitingForServiceNameTransition() []statemachine.Transition {
	return []statemachine.Transition{
		{CustomMessage, &statemachine.Transaction{
			CurrentState: WaitingForPasswordExist,
			Transitions:  waitingForPasswordExistTransition(),
		}},
	}
}

func waitingForPasswordExistTransition() []statemachine.Transition {
	return []statemachine.Transition{
		{YesCommand, &statemachine.Transaction{
			CurrentState: WaitingForPassword,
			Transitions:  waitingForPasswordTransition(),
		}},
		{NoCommand, &statemachine.Transaction{
			CurrentState: ChoiceCharactersNumber,
			Transitions:  choiceCharactersNumberTransaction(),
		}},
	}
}

func choiceCharactersNumberTransaction() []statemachine.Transition {
	return []statemachine.Transition{
		{CustomMessage, &statemachine.Transaction{
			CurrentState: ChoiceCharactersAlphabet,
			Transitions:  nil,
		}},
	}
}

func waitingForPasswordTransition() []statemachine.Transition {
	return []statemachine.Transition{
		{
			CustomMessage, nil,
		},
	}
}
