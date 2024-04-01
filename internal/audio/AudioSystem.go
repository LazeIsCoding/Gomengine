package audio

import (
	msgb "Gomengine/internal/messagebus"
	"Gomengine/pkg/datastructures"
	"fmt"
)

type AudioSystem struct {
	msgb.System
}

func NewAudioSystem() *AudioSystem {
	return &AudioSystem{
		System: msgb.System{
			Name:      "Audio",
			MsgQue:    *datastructures.NewQueue[msgb.Msg](),
			ImmMsgQue: *datastructures.NewQueue[msgb.Msg](),
		},
	}
}

func (s *AudioSystem) HandleMsg() {
	if !s.ImmMsgQue.IsEmpty() {

	} else {
		msg := s.MsgQue.Dequeue()
		switch msg.Message {
		case msgb.EXAMPLE_MESSAGE:
			fmt.Print("Handle Message from AudioSystem\n")
		}
	}
}
