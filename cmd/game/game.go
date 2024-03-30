package main

import (
	"Gomengine/internal/audio"
	"Gomengine/internal/graphics"
	"Gomengine/internal/gui"
	"Gomengine/internal/messagebus"
	"Gomengine/internal/physics"
	"fmt"
)

func main() {

	var msgBus = messagebus.InitNewBus()
	var graphSys = &graphics.GraphicsSystem{}
	var audSys = &audio.AudioSystem{}
	var guiSys = &gui.GuiSystem{}
	var phySys = &physics.PhysicsSystem{}

	msgBus.RegisterSystem("Graphics", graphSys)
	msgBus.RegisterSystem("Audio", audSys)
	msgBus.RegisterSystem("Gui", guiSys)
	msgBus.RegisterSystem("Physics", phySys)

	msgBus.BroadCast(messagebus.Msg{Message: messagebus.EXAMPLE_MESSAGE})
	for _, system := range msgBus.Systems {
		system.HandleMsg()
	}
	fmt.Print("finished")
}
