package commands

type ServerHudSetParam struct{}

func (p *ServerHudSetParam) GetCommandId() uint16 {
	return ServerCommandHudSetParam
}

func (p *ServerHudSetParam) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerHudSetParam) UnmarshalPacket(payload []byte) error {
	return nil
}

func (p *ServerHudSetParam) String() string {
	return "{ServerHudSetParam}"
}
