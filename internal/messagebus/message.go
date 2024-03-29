package messagebus

type Message int

const (
	EXAMPLE_MESSAGE Message = 0
)

type Msg struct {
	Message
}
