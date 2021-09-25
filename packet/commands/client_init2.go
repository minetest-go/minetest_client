package commands

type ClientInit2 struct {
}

func NewClientInit2() *ClientInit2 {
	return &ClientInit2{}
}

func (p *ClientInit2) GetCommandId() uint16 {
	return ClientCommandInit2
}

func (p *ClientInit2) MarshalPacket() ([]byte, error) {
	return []byte{}, nil
}

func (p *ClientInit2) UnmarshalPacket([]byte) error {
	return nil
}

func (p *ClientInit2) String() string {
	return "{ClientInit2}"
}
