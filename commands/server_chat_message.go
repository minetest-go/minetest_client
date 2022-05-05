package commands

import (
	"encoding/binary"
	"fmt"

	"golang.org/x/text/encoding/unicode"
)

type ServerChatMessageType int

type ServerChatMessage struct {
	Version int
	Type    ServerChatMessageType
	Sender  string
	Message string
}

func (p *ServerChatMessage) GetCommandId() uint16 {
	return ServerCommandChatMessage
}

func (p *ServerChatMessage) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerChatMessage) UnmarshalPacket(payload []byte) error {
	offset := 0
	p.Version = int(payload[offset])
	offset++
	p.Type = ServerChatMessageType(payload[offset])
	offset++

	sender_len := binary.BigEndian.Uint16(payload[offset:])
	offset += 2

	utf16 := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	sender_bytes, err := utf16.NewDecoder().Bytes(payload[offset : offset+int(sender_len*2)])
	if err != nil {
		return err
	}

	p.Sender = string(sender_bytes)
	offset += int(sender_len * 2)

	message_len := binary.BigEndian.Uint16(payload[offset:])
	offset += 2

	message_bytes, err := utf16.NewDecoder().Bytes(payload[offset : offset+int(message_len*2)])
	if err != nil {
		return err
	}

	p.Message = string(message_bytes)

	return nil
}

func (p *ServerChatMessage) String() string {
	return fmt.Sprintf("{ServerChatMessage Version=%d, Type=%d, Sender='%s', Message='%s'}",
		p.Version, p.Type, p.Sender, p.Message)
}
