package commands

import (
	"encoding/binary"
	"fmt"
)

type ServerSRPBytesSB struct {
	BytesS []byte
	BytesB []byte
}

func (p *ServerSRPBytesSB) GetCommandId() uint16 {
	return ServerCommandSRPBytesSB
}

func (p *ServerSRPBytesSB) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerSRPBytesSB) UnmarshalPacket(payload []byte) error {
	bytes_s_length := binary.BigEndian.Uint16(payload[0:])
	p.BytesS = payload[2 : bytes_s_length+2]
	bytes_b_length := binary.BigEndian.Uint16(payload[:bytes_s_length+2])
	p.BytesB = payload[bytes_s_length+2+2 : bytes_b_length+bytes_s_length+2+2]
	return nil
}

func (p *ServerSRPBytesSB) String() string {
	return fmt.Sprintf("{ServerSRPBytesSB #BytesS=%d, #BytesB=%d}",
		len(p.BytesS), len(p.BytesB))
}
