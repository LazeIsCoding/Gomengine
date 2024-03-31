package messagebus

type Message int

const (
	EXAMPLE_MESSAGE Message = 0
	TEST_MESSAGE    Message = 1
)

type Msg struct {
	Message
	Immediate bool
}
