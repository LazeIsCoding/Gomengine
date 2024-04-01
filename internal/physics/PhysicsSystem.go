package physics

import (
	msgb "Gomengine/internal/messagebus"
	"Gomengine/pkg/datastructures"
	"fmt"
)

type PhysicsSystem struct {
	msgb.System
}

func NewPhysicsSystem() *PhysicsSystem {
	return &PhysicsSystem{
		System: msgb.System{
			Name:      "Physics",
			MsgQue:    *datastructures.NewQueue[msgb.Msg](),
			ImmMsgQue: *datastructures.NewQueue[msgb.Msg](),
		},
	}
}

func (s *PhysicsSystem) HandleMsg() {
	if !s.ImmMsgQue.IsEmpty() {

	} else {
		msg := s.MsgQue.Dequeue()
		switch msg.Message {
		case msgb.EXAMPLE_MESSAGE:
			fmt.Print("Handle Message from PhysicsSystem\n")
		}
	}
}
