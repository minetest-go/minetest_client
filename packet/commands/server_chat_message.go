package commands

import (
	"fmt"
)

type ServerChatMessage struct {
	Message string
}

func (p *ServerChatMessage) GetCommandId() uint16 {
	return 47
}

func (p *ServerChatMessage) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerChatMessage) UnmarshalPacket(payload []byte) error {
	//fmt.Printf("Chat message: len=%d %s\n", len(payload), fmt.Sprint(payload))
	size := payload[5]
	p.Message = string(payload[6 : (size*2)+6])
	return nil
}

func (p *ServerChatMessage) String() string {
	return fmt.Sprintf("{ServerChatMessage Message=%s}", p.Message)
}
