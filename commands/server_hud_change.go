package commands

type ServerHudChange struct{}

func (p *ServerHudChange) GetCommandId() uint16 {
	return ServerCommandHudChange
}

func (p *ServerHudChange) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerHudChange) UnmarshalPacket(payload []byte) error {
	return nil
}

func (p *ServerHudChange) String() string {
	return "{ServerHudChange}"
}
