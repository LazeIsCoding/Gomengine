package main

import (
	"Gomengine/internal/audio"
	"Gomengine/internal/graphics"
	"Gomengine/internal/gui"
	"Gomengine/internal/messagebus"
	"Gomengine/internal/physics"
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("Fehler beim Lesen der Eingabe: ", err)
			return
		}

		input := scanner.Text()
		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Error")
			continue
		}
		msg := messagebus.Msg{
			Message:   messagebus.Message(num),
			Immediate: false,
		}

		msgBus.BroadCast(msg)
		for _, system := range msgBus.Systems {
			system.HandleMsg()
		}
	}

	fmt.Print("finished")
}
