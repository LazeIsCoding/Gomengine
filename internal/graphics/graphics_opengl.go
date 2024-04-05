package graphics

import (
	msgb "Gomengine/internal/messagebus"
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	"runtime"
)

/*
func InitOpenGL() error {

}*/

func CreateWindow(width, height int, windowName string, fullscreen bool, xPos, yPos int) error {
	runtime.LockOSThread()
	if err := glfw.Init(); err != nil {
		return err
	}
	defer glfw.Terminate()
	window, err := glfw.CreateWindow(width, height, windowName, nil, nil)
	if err != nil {
		return err
	}
	fmt.Println("Successfully opened window with Width:", width, "px, Height:", height, "px")
	window.MakeContextCurrent()

	window.SetPos(xPos, yPos)
	for !window.ShouldClose() {
		//OPENGL
		window.SwapBuffers()
		glfw.PollEvents()
	}
	defer window.Destroy()
	return nil

}

func HandleOpenGLMessage(msg msgb.Msg) {

	switch msg.Message {
	case msgb.CREATE_WINDOW:
		fmt.Println("Handle Create Window from OpenGL Graphics")
		windowData := msg.Data.(msgb.WindowData)
		width := windowData.Width
		height := windowData.Height
		name := windowData.Name
		go CreateWindow(width, height, name, false, 200, 200)

		break
	}

}
