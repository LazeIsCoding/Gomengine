package main

import (
	"Gomengine/internal/gomengine"
)

func main() {

	update()

}

func init() {
	gomengine.InitGame()
	gomengine.CreateWindow(200, 500)
}

func update() {
	gomengine.SendMessage()
}
