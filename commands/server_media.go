package commands

import (
	"encoding/binary"
	"fmt"
)

type ServerMedia struct {
	Bunches  uint16
	Index    uint16
	NumFiles uint32
	Files    map[string][]byte
}

func (p *ServerMedia) GetCommandId() uint16 {
	return ServerCommandMedia
}

func (p *ServerMedia) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerMedia) UnmarshalPacket(payload []byte) error {
	offset := 0
	p.Bunches = binary.BigEndian.Uint16(payload[offset:])
	offset += 2
	p.Index = binary.BigEndian.Uint16(payload[offset:])
	offset += 2
	p.NumFiles = binary.BigEndian.Uint32(payload[offset:])
	offset += 4

	p.Files = make(map[string][]byte)

	for i := 0; i < int(p.NumFiles); i++ {
		name_len := binary.BigEndian.Uint16(payload[offset:])
		offset += 2

		name := string(payload[offset : offset+int(name_len)])
		offset += int(name_len)

		data_len := binary.BigEndian.Uint32(payload[offset:])
		offset += 4

		p.Files[name] = payload[offset : offset+int(data_len)]
		offset += int(data_len)
	}

	return nil
}

func (p *ServerMedia) String() string {
	return fmt.Sprintf("{ServerMedia Bunches=%d, Index=%d, NumFiles=%d}", p.Bunches, p.Index, p.NumFiles)
}
