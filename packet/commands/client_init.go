package commands

import "encoding/binary"

type ClientInit struct {
	ClientMax                 int
	SupportedCompressionModes uint16
	MinNetProtoVersion        uint16
	MaxNetProtoVersion        uint16
	PlayerName                string
}

func NewClientInit(playername string) *ClientInit {
	return &ClientInit{
		ClientMax:                 28,
		SupportedCompressionModes: 0,
		MinNetProtoVersion:        37,
		MaxNetProtoVersion:        39,
		PlayerName:                playername,
	}
}

func (p *ClientInit) GetCommandId() uint16 {
	return 2
}

func (p *ClientInit) MarshalPacket() ([]byte, error) {
	packet := make([]byte, 1+2+2+2)
	packet[0] = byte(p.ClientMax)
	binary.BigEndian.PutUint16(packet[1:], p.SupportedCompressionModes)
	binary.BigEndian.PutUint16(packet[3:], p.MinNetProtoVersion)
	binary.BigEndian.PutUint16(packet[5:], p.MaxNetProtoVersion)

	packet = append(packet, 0, byte(len(p.PlayerName)))
	packet = append(packet, []byte(p.PlayerName)...)

	return packet, nil
}

func (p *ClientInit) UnmarshalPacket([]byte) error {
	return nil
}
