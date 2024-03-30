package graphics

import (
	msgb "Gomengine/internal/messagebus"
	"fmt"
)

type GraphicsSystem struct {
	msgb.System
}

func (s *GraphicsSystem) HandleMsg() {
	msg := s.MsgQue[0]
	switch msg.Message {
	case msgb.EXAMPLE_MESSAGE:
		fmt.Print("Handle Message from GraphicsSystem\n")
	}
}
