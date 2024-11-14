package commands

import (
	"encoding/binary"
	"fmt"
)

type ServerSetPeer struct {
	PeerID uint16
}

func (p *ServerSetPeer) GetCommandId() uint16 {
	return ServerCommandSetPeer
}

func (p *ServerSetPeer) MarshalPacket() ([]byte, error) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, p.PeerID)
	return buf, nil
}

func (p *ServerSetPeer) UnmarshalPacket(payload []byte) error {
	p.PeerID = binary.BigEndian.Uint16(payload)
	return nil
}

func (p *ServerSetPeer) String() string {
	return fmt.Sprintf("{ServerSetPeer PeerID=%d}", p.PeerID)
}
