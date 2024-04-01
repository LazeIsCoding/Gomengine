package gui

import (
	msgb "Gomengine/internal/messagebus"
	"Gomengine/pkg/datastructures"
	"fmt"
)

type GuiSystem struct {
	msgb.System
}

func NewGuiSystem() *GuiSystem {
	return &GuiSystem{
		System: msgb.System{
			Name:      "Gui",
			MsgQue:    *datastructures.NewQueue[msgb.Msg](),
			ImmMsgQue: *datastructures.NewQueue[msgb.Msg](),
		},
	}
}

func (s *GuiSystem) HandleMsg() {
	if !s.ImmMsgQue.IsEmpty() {

	} else {
		msg := s.MsgQue.Dequeue()
		switch msg.Message {
		case msgb.EXAMPLE_MESSAGE:
			fmt.Print("Handle Message from GuiSystem\n")
		}
	}
}
