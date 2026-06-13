package commands

import (
	"encoding/binary"
	"fmt"

	"golang.org/x/text/encoding/unicode"
)

type ClientChatMessage struct {
	Message string
}

func NewClientChatMessage(msg string) *ClientChatMessage {
	return &ClientChatMessage{Message: msg}
}

func (p *ClientChatMessage) GetCommandId() uint16 {
	return ClientCommandChatMessage
}

func (p *ClientChatMessage) MarshalPacket() ([]byte, error) {
	utf16be := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	encoder := utf16be.NewEncoder()

	encoded, err := encoder.Bytes([]byte(p.Message))
	if err != nil {
		return nil, err
	}

	charCount := uint16(len(encoded) / 2)
	data := make([]byte, 2+len(encoded))
	binary.BigEndian.PutUint16(data[0:2], charCount)
	copy(data[2:], encoded)

	return data, nil
}

func (p *ClientChatMessage) UnmarshalPacket([]byte) error {
	return nil
}

func (p *ClientChatMessage) String() string {
	return fmt.Sprintf("{ClientChatMessage Message='%s'}", p.Message)
}
