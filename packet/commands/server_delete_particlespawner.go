package commands

import (
	"encoding/binary"
	"fmt"
)

type ServerDeleteParticleSpawner struct {
	ServerID uint32
}

func (p *ServerDeleteParticleSpawner) GetCommandId() uint16 {
	return 83
}

func (p *ServerDeleteParticleSpawner) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerDeleteParticleSpawner) UnmarshalPacket(payload []byte) error {
	p.ServerID = binary.BigEndian.Uint32(payload)
	return nil
}

func (p *ServerDeleteParticleSpawner) String() string {
	return fmt.Sprintf("{ServerDeleteParticleSpawner ServerID=%d}", p.ServerID)
}
