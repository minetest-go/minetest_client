package commands

import "encoding/binary"

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
	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, uint16(len(p.Salt)))
	data = append(data, p.Salt...)

	vkey_len_bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(vkey_len_bytes, uint16(len(p.VerificationKey)))
	data = append(data, vkey_len_bytes...)
	data = append(data, p.VerificationKey...)
	data = append(data, 0)

	return data, nil
}

func (p *ClientFirstSRP) UnmarshalPacket([]byte) error {
	return nil
}

func (p *ClientFirstSRP) String() string {
	return "{ClientFirstSRP}"
}
