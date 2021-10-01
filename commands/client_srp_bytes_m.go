package commands

import (
	"encoding/binary"
	"fmt"
)

type ClientSRPBytesM struct {
	BytesM []byte
}

func NewClientSRPBytesM(bytes_m []byte) *ClientSRPBytesM {
	return &ClientSRPBytesM{
		BytesM: bytes_m,
	}
}

func (p *ClientSRPBytesM) GetCommandId() uint16 {
	return ClientCommandSRPBytesM
}

func (p *ClientSRPBytesM) MarshalPacket() ([]byte, error) {
	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, uint16(len(p.BytesM)))
	data = append(data, p.BytesM...)
	data = append(data, 0x01)

	return data, nil
}

func (p *ClientSRPBytesM) UnmarshalPacket([]byte) error {
	return nil
}

func (p *ClientSRPBytesM) String() string {
	return fmt.Sprintf("{ClientSRPBytesM #BytesM=%d}", len(p.BytesM))
}
