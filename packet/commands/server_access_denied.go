package commands

type ServerAccessDenied struct{}

func (p *ServerAccessDenied) GetCommandId() uint16 {
	return 10
}

func (p *ServerAccessDenied) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerAccessDenied) UnmarshalPacket(payload []byte) error {
	return nil
}

func (p *ServerAccessDenied) String() string {
	return "{ServerAccessDenied}"
}
