package main

import (
	"Gomengine/internal/gomengine"
)

func main() {

	update()

}

func init() {
	gomengine.InitGame()
	gomengine.CreateWindow(1000, 500, "FirstWindow")
}

func update() {
	gomengine.SendMessage()
}
