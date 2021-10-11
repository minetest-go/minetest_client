package commands

import (
	"encoding/binary"
	"fmt"
)

type ServerAnnounceMedia struct {
	FileCount uint16
}

func (p *ServerAnnounceMedia) GetCommandId() uint16 {
	return ServerCommandAnnounceMedia
}

func (p *ServerAnnounceMedia) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerAnnounceMedia) UnmarshalPacket(payload []byte) error {
	p.FileCount = binary.BigEndian.Uint16(payload[0:])

	return nil
}

func (p *ServerAnnounceMedia) String() string {
	return fmt.Sprintf("{ServerAnnounceMedia FileCount=%d}", p.FileCount)
}
