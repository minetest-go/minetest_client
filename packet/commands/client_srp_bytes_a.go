package commands

import (
	"encoding/binary"
	"fmt"
)

type ClientSRPBytesA struct {
	BytesA []byte
}

func NewClientSRPBytesA(bytes_a []byte) *ClientSRPBytesA {
	return &ClientSRPBytesA{
		BytesA: bytes_a,
	}
}

func (p *ClientSRPBytesA) GetCommandId() uint16 {
	return ClientCommandSRPBytesA
}

func (p *ClientSRPBytesA) MarshalPacket() ([]byte, error) {
	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, uint16(len(p.BytesA)))
	data = append(data, p.BytesA...)
	data = append(data, 0x01)

	return data, nil
}

func (p *ClientSRPBytesA) UnmarshalPacket([]byte) error {
	return nil
}

func (p *ClientSRPBytesA) String() string {
	return fmt.Sprintf("{ClientSRPBytesA #BytesA=%d}", len(p.BytesA))
}
