package commands

type ServerHudAdd struct{}

func (p *ServerHudAdd) GetCommandId() uint16 {
	return ServerCommandHudAdd
}

func (p *ServerHudAdd) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerHudAdd) UnmarshalPacket(payload []byte) error {
	return nil
}

func (p *ServerHudAdd) String() string {
	return "{ServerHudAdd}"
}
