package messagebus

type Message int

const (
	EXAMPLE_MESSAGE Message = 0
	TEST_MESSAGE    Message = 1
)

var (
	immediate bool
)

type Msg struct {
	Message
}
