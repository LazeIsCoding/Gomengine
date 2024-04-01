package gomengine

import (
	"Gomengine/internal/audio"
	"Gomengine/internal/graphics"
	"Gomengine/internal/gui"
	msgb "Gomengine/internal/messagebus"
	"Gomengine/internal/physics"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var gomen gomengine

type gomengine struct {
	MsgbBus msgb.MessageBus
}

func InitGame() {
	gomen.MsgbBus = *msgb.InitNewBus()
	var graphSys = graphics.NewGraphicsSystem()
	var audSys = audio.NewAudioSystem()
	var guiSys = gui.NewGuiSystem()
	var phySys = physics.NewPhysicsSystem()

	gomen.MsgbBus.RegisterSystem(graphSys.Name, graphSys)
	gomen.MsgbBus.RegisterSystem(audSys.Name, audSys)
	gomen.MsgbBus.RegisterSystem(guiSys.Name, guiSys)
	gomen.MsgbBus.RegisterSystem(phySys.Name, phySys)

	/*for _, system := range gomen.MsgbBus.Systems {
		go system.Update()
	}*/
	go audSys.Update()
	go phySys.Update()
	go graphSys.Update()
	go guiSys.Update()

}

func SendMessage() {
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

		msg := msgb.Msg{
			Message:   msgb.Message(num),
			Immediate: false,
		}
		gomen.MsgbBus.BroadCast(msg)

	}

}

func CreateWindow(width, height int) {
	msg := msgb.Msg{
		Message:   msgb.CREATE_WINDOW,
		Immediate: false,
		Data:      msgb.WindowData{Width: width, Height: height},
	}
	gomen.MsgbBus.BroadCast(msg)
}
