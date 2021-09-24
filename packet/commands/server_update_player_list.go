package commands

import (
	"fmt"
)

type PlayerListModifier int

const (
	Added PlayerListModifier = iota
	Removed
	Init
)

type ServerUpdatePlayerList struct {
	Players  []string
	Modifier PlayerListModifier
}

func (p *ServerUpdatePlayerList) GetCommandId() uint16 {
	return 86
}

func (p *ServerUpdatePlayerList) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerUpdatePlayerList) UnmarshalPacket(payload []byte) error {
	fmt.Printf("ServerUpdatePlayerList: len=%d %s\n", len(payload), fmt.Sprint(payload))
	// TODO
	return nil
}

func (p *ServerUpdatePlayerList) String() string {
	return fmt.Sprintf("{ServerChatMessage Players='%s'}", p.Players)
}
