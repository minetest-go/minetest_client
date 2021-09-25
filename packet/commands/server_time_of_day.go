package commands

import (
	"encoding/binary"
	"fmt"
)

type ServerTimeOfDay struct {
	TimeOfDay int
}

func (p *ServerTimeOfDay) GetCommandId() uint16 {
	return ServerCommandTimeOfDay
}

func (p *ServerTimeOfDay) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerTimeOfDay) UnmarshalPacket(payload []byte) error {
	tod_raw := binary.BigEndian.Uint16(payload)
	p.TimeOfDay = int(tod_raw) % 24000
	return nil
}

func (p *ServerTimeOfDay) String() string {
	return fmt.Sprintf("{ServerTimeOfDay TimeOfDay=%d}", p.TimeOfDay)
}
