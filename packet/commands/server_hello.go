package commands

import (
	"encoding/binary"
	"fmt"
)

type ServerHello struct {
	SerializationVersion uint8
	CompressionMode      uint16
	ProtocolVersion      uint16
}

func (p *ServerHello) GetCommandId() uint16 {
	return 2
}

func (p *ServerHello) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerHello) UnmarshalPacket(payload []byte) error {
	p.SerializationVersion = payload[0]
	p.CompressionMode = binary.BigEndian.Uint16(payload[1:])
	p.ProtocolVersion = binary.BigEndian.Uint16(payload[3:])
	return nil
}

func (p *ServerHello) String() string {
	return fmt.Sprintf("{ServerHello SerializationVersion=%d, CompressionMode=%d, ProtocolVersion=%d}",
		p.SerializationVersion, p.CompressionMode, p.ProtocolVersion)
}
