package commands

type ServerHudSetFlags struct{}

func (p *ServerHudSetFlags) GetCommandId() uint16 {
	return ServerCommandHudSetFlags
}

func (p *ServerHudSetFlags) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerHudSetFlags) UnmarshalPacket(payload []byte) error {
	return nil
}

func (p *ServerHudSetFlags) String() string {
	return "{ServerHudSetFlags}"
}
