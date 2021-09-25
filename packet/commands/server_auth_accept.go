package commands

import (
	"encoding/binary"
	"fmt"
	"math"
)

type ServerAuthAccept struct {
	PosX         uint32
	PosY         uint32
	PosZ         uint32
	Seed         uint64
	SendInterval float32
}

func (p *ServerAuthAccept) GetCommandId() uint16 {
	return ServerCommandAuthAccept
}

func (p *ServerAuthAccept) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerAuthAccept) UnmarshalPacket(payload []byte) error {
	p.PosX = binary.BigEndian.Uint32(payload[0:])
	p.PosY = binary.BigEndian.Uint32(payload[4:])
	p.PosZ = binary.BigEndian.Uint32(payload[8:])
	p.Seed = binary.BigEndian.Uint64(payload[12:])
	bits := binary.BigEndian.Uint32(payload[20:])
	p.SendInterval = math.Float32frombits(bits)
	return nil
}

func (p *ServerAuthAccept) String() string {
	return fmt.Sprintf("{ServerAuthAccept PosX=%d, PosY=%d, PosZ=%d, Seed=%d, SendInterval=%f}",
		p.PosX, p.PosY, p.PosZ, p.Seed, p.SendInterval)
}
