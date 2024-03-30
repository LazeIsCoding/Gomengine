package physics

import (
	msgb "Gomengine/internal/messagebus"
	"fmt"
)

type PhysicsSystem struct {
	msgb.System
}

func (s *PhysicsSystem) HandleMsg() {
	msg := s.MsgQue[0]
	switch msg.Message {
	case msgb.EXAMPLE_MESSAGE:
		fmt.Print("Handle Message from PhysicsSystem\n")
	}
}
