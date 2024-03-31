package gui

import (
	msgb "Gomengine/internal/messagebus"
	"fmt"
)

type GuiSystem struct {
	msgb.System
}

func (s *GuiSystem) HandleMsg() {
	msg := s.MsgQue[0]
	switch msg.Message {
	case msgb.EXAMPLE_MESSAGE:
		fmt.Print("Handle Example Message from GuiSystem\n")
	}
}
