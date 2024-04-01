package graphics

import (
	msgb "Gomengine/internal/messagebus"
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func InitOpenGL() error {

	if err := glfw.Init(); err != nil {
		return err
	}
	defer glfw.Terminate()

	return nil
}

func HandleOpenGLMessage(msg msgb.Msg) {

	switch msg.Message {
	case msgb.CREATE_WINDOW:
		windowData := msg.Data.(msgb.WindowData)
		width := windowData.Width
		height := windowData.Height
		fmt.Println("Graphics should open a window with Height:", height, " Width:", width)
		break
	}

}
