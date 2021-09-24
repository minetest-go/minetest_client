package commands

import "encoding/binary"

type ServerSetPeer struct {
	PeerID uint16
}

func (p *ServerSetPeer) GetCommandId() uint8 {
	return 1
}

func (p *ServerSetPeer) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerSetPeer) UnmarshalPacket(payload []byte) error {
	p.PeerID = binary.BigEndian.Uint16(payload)
	return nil
}
