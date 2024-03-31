package messagebus

import "fmt"

type messageBus struct {
	Systems map[string]SystemInterface
}

func InitNewBus() *messageBus {
	return &messageBus{Systems: make(map[string]SystemInterface)}
}

// Register a System to the Message Bus
func (mb *messageBus) RegisterSystem(name string, sys SystemInterface) {
	mb.Systems[name] = sys
}

// Broadcast a message to all registered Systems
func (m *messageBus) BroadCast(msg Msg) {
	for _, system := range m.Systems {
		system.RecMsg(msg)
	}
}

type SystemInterface interface {
	Equals(other SystemInterface) bool
	RecMsg(msg Msg)
	ShowMsgs()
	HandleMsg()
}

type System struct {
	name      string
	MsgQue    []Msg
	ImmMsgQue []Msg
}

func (s System) Equals(other SystemInterface) bool {
	return s.name == other.(*System).name
}

func (s *System) RecMsg(msg Msg) {
	if msg.Immediate == true {
		s.ImmMsgQue = append(s.MsgQue, msg)
	}
	s.MsgQue = append(s.MsgQue, msg)
}

func (s *System) ShowMsgs() {
	for i, msg := range s.MsgQue {
		fmt.Printf("Message: %d, %d", i, msg.Message)
	}
}

func (s *System) HandleMsg() {
	panic("Each System should implement handleMsg independently")
}
