package commands

type ClientPeerInit struct {
}

func NewClientPeerInit() *ClientPeerInit {
	return &ClientPeerInit{}
}

func (p *ClientPeerInit) GetCommandId() uint16 {
	return 0
}

func (p *ClientPeerInit) MarshalPacket() ([]byte, error) {
	return []byte{0, 0}, nil
}

func (p *ClientPeerInit) UnmarshalPacket([]byte) error {
	return nil
}
