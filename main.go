package main

import (
	"github.com/ent1k1377/statemachine/dialogues"
	sm "github.com/ent1k1377/statemachine/statemachine"
)

func main() {
	statemachine := sm.New(dialogues.InitialState)
	dialogues.NewInitialDialogue(statemachine)
	dialogues.NewServiceDialogue()
	statemachine.PerformTransition(dialogues.StartCommand)
}
