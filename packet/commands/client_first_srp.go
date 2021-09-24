package commands

type ClientFirstSRP struct {
	Salt            []byte
	VerificationKey []byte
}

func NewClientFirstSRP(salt []byte, verificationKey []byte) *ClientFirstSRP {
	return &ClientFirstSRP{
		Salt:            salt,
		VerificationKey: verificationKey,
	}
}

func (p *ClientFirstSRP) GetCommandId() uint16 {
	return 80
}

func (p *ClientFirstSRP) MarshalPacket() ([]byte, error) {
	return []byte{0, 0}, nil
}

func (p *ClientFirstSRP) UnmarshalPacket([]byte) error {
	return nil
}

func (p *ClientFirstSRP) String() string {
	return "{ClientFirstSRP}"
}
