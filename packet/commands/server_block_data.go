package commands

import (
	"encoding/binary"
	"fmt"
)

type ServerBlockData struct {
	PosX uint16
	PosY uint16
	PosZ uint16
}

func (p *ServerBlockData) GetCommandId() uint16 {
	return ServerCommandBlockData
}

func (p *ServerBlockData) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerBlockData) UnmarshalPacket(payload []byte) error {
	p.PosX = binary.BigEndian.Uint16(payload[0:])
	p.PosY = binary.BigEndian.Uint16(payload[2:])
	p.PosZ = binary.BigEndian.Uint16(payload[4:])
	return nil
}

func (p *ServerBlockData) String() string {
	return fmt.Sprintf("{ServerBlockData pos=%d/%d%d}",
		p.PosX, p.PosY, p.PosZ)
}
