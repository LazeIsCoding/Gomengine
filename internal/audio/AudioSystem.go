package audio

import (
	msgb "Gomengine/internal/messagebus"
	"fmt"
)

type AudioSystem struct {
	msgb.System
}

func (s *AudioSystem) HandleMsg() {
	msg := s.MsgQue[0]
	switch msg.Message {
	case msgb.EXAMPLE_MESSAGE:
		fmt.Print("Handle Message from AudioSystem\n")
	}
}
