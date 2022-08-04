package msg

import "fmt"

type Message interface {
	GetMessageType() uint32
}

const (
	_ uint32 = iota

	PingMessageType
	PoneMessageType

	NameMessageType
	NameResultMessageType
)

func NewMessage(messageType uint32) (message Message, err error) {
	switch messageType {
	case PingMessageType:
		message = &PingMessage{}
	case PoneMessageType:
		message = &PoneMessage{}

	case NameMessageType:
		message = &NameMessage{}
	case NameResultMessageType:
		message = &NameResultMessage{}
	default:
		err = fmt.Errorf("no find message type: %d", messageType)
	}
	return
}
