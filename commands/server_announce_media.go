package commands

type ServerAnnounceMedia struct{}

func (p *ServerAnnounceMedia) GetCommandId() uint16 {
	return ServerCommandAnnounceMedia
}

func (p *ServerAnnounceMedia) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerAnnounceMedia) UnmarshalPacket(payload []byte) error {
	return nil
}

func (p *ServerAnnounceMedia) String() string {
	return "{ServerAnnounceMedia}"
}
