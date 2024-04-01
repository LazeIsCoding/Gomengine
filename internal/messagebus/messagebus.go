package messagebus

import (
	"Gomengine/pkg/datastructures"
	"fmt"
)

type MessageBus struct {
	Systems map[string]SystemInterface
}

func InitNewBus() *MessageBus {
	return &MessageBus{Systems: make(map[string]SystemInterface)}
}

// Register a System to the Message Bus
func (mb *MessageBus) RegisterSystem(name string, sys SystemInterface) {
	fmt.Println("Registered System: ", name)
	mb.Systems[name] = sys
}

// Broadcast a message to all registered Systems
func (m *MessageBus) BroadCast(msg Msg) {
	fmt.Println("Broadcast of Message: ", Message(msg.Message))
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
	Name      string
	MsgQue    datastructures.Queue[Msg]
	ImmMsgQue datastructures.Queue[Msg]
}

func (s System) Equals(other SystemInterface) bool {
	return s.Name == other.(*System).Name
}

func (s *System) RecMsg(msg Msg) {
	fmt.Println("Message received: ", msg.Message, "from: ", s.Name)
	if msg.Immediate == true {
		s.ImmMsgQue.Enqueue(msg)
	} else {
		s.MsgQue.Enqueue(msg)
	}
}

func (s *System) ShowMsgs() {
	s.MsgQue.PrintQueue()
}

func (s *System) HandleMsg() {
	panic("Each System should implement handleMsg independently")
}
