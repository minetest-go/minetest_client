package commands

type ServerMovement struct{}

func (p *ServerMovement) GetCommandId() uint16 {
	return ServerCommandMovement
}

func (p *ServerMovement) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerMovement) UnmarshalPacket(payload []byte) error {
	return nil
}

func (p *ServerMovement) String() string {
	return "{ServerMovement}"
}
