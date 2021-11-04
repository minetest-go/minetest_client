package commands

import (
	"encoding/binary"
	"fmt"
	"math"
)

type ServerMovePlayer struct {
	X     float32
	Y     float32
	Z     float32
	Pitch float32
	Yaw   float32
}

func (p *ServerMovePlayer) GetCommandId() uint16 {
	return ServerCommandMovePlayer
}

func (p *ServerMovePlayer) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerMovePlayer) UnmarshalPacket(payload []byte) error {
	p.X = math.Float32frombits(binary.BigEndian.Uint32(payload[0:]))
	p.Y = math.Float32frombits(binary.BigEndian.Uint32(payload[4:]))
	p.Z = math.Float32frombits(binary.BigEndian.Uint32(payload[8:]))
	p.Pitch = math.Float32frombits(binary.BigEndian.Uint32(payload[12:]))
	p.Yaw = math.Float32frombits(binary.BigEndian.Uint32(payload[16:]))
	return nil
}

func (p *ServerMovePlayer) String() string {
	return fmt.Sprintf("{ServerMovePlayer pos=%f/%f/%f, pitch=%f, yaw=%f}", p.X, p.Y, p.Z, p.Pitch, p.Yaw)
}
