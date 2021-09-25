package commands

type ClientPeerInit struct {
}

func NewClientPeerInit() *ClientPeerInit {
	return &ClientPeerInit{}
}

func (p *ClientPeerInit) GetCommandId() uint16 {
	return ClientCommandPeerInit
}

func (p *ClientPeerInit) MarshalPacket() ([]byte, error) {
	return []byte{}, nil
}

func (p *ClientPeerInit) UnmarshalPacket([]byte) error {
	return nil
}

func (p *ClientPeerInit) String() string {
	return "{ClientPeerInit}"
}
