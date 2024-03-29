package gui

import (
	msgb "Gomengine/internal/messagebus"
)

type GuiSystem struct {
}

func (s *GuiSystem) handleMessage(msg msgb.Msg) {
	switch msg.Message {
	case msgb.EXAMPLE_MESSAGE:
		break
	}
}
