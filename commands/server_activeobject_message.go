package commands

type ServerActiveObjectMessage struct{}

func (p *ServerActiveObjectMessage) GetCommandId() uint16 {
	return ServerCommandActiveObjectMessage
}

func (p *ServerActiveObjectMessage) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerActiveObjectMessage) UnmarshalPacket(payload []byte) error {
	return nil
}

func (p *ServerActiveObjectMessage) String() string {
	return "{ServerActiveObjectMessage}"
}
