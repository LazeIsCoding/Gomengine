package messagebus

type Message int

const (
	EXAMPLE_MESSAGE Message = 0
	TEST_MESSAGE    Message = 1
	CREATE_WINDOW   Message = 2
)

type WindowData struct {
	Width  int
	Height int
}

type Msg struct {
	Message
	Immediate bool
	Data      interface{}
}
