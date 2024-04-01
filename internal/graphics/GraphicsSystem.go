package graphics

import (
	msgb "Gomengine/internal/messagebus"
	"Gomengine/pkg/datastructures"
	"fmt"
)

type GraphicsSystem struct {
	msgb.System
}

func NewGraphicsSystem() *GraphicsSystem {
	return &GraphicsSystem{
		System: msgb.System{
			Name:      "Graphics",
			MsgQue:    *datastructures.NewQueue[msgb.Msg](),
			ImmMsgQue: *datastructures.NewQueue[msgb.Msg](),
		},
	}
}

func (s *GraphicsSystem) HandleMsg() {
	if !s.ImmMsgQue.IsEmpty() {

	} else {
		msg := s.MsgQue.Dequeue()
		switch msg.Message {
		case msgb.EXAMPLE_MESSAGE:
			fmt.Print("Handle Message from GraphicSystem\n")
		}
	}
}
