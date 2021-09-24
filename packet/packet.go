package packet

import "encoding/binary"

type Packet struct {
	Command    Command
	PacketType PacketType
	SubType    PacketType
	PeerID     uint16
	SeqNr      uint16
	Channel    uint8
}

func CreateReliable(peerId uint16, seqNr uint16, command Command) *Packet {
	return &Packet{
		PacketType: Reliable,
		SubType:    Original,
		Command:    command,
		PeerID:     peerId,
		SeqNr:      seqNr,
	}
}

func CreatePacket(packetType PacketType, subType PacketType, peerId uint16, seqNr uint16, command Command) *Packet {
	return &Packet{
		PacketType: packetType,
		SubType:    subType,
		Command:    command,
		PeerID:     peerId,
		SeqNr:      seqNr,
	}
}

func (p *Packet) MarshalPacket() ([]byte, error) {
	packet := make([]byte, 4+2+1+1)
	copy(packet[0:], ProtocolID)
	binary.BigEndian.PutUint16(packet[4:], p.PeerID)
	packet[6] = p.Channel
	packet[7] = p.Command.GetCommandId()

	if p.PacketType == Reliable {
		bytes := make([]byte, 2)
		binary.BigEndian.PutUint16(bytes, p.SeqNr)
		packet = append(packet, bytes...)
		packet = append(packet, byte(p.SubType))
	}

	payload, err := p.Command.MarshalPacket()
	packet = append(packet, payload...)
	return packet, err
}

func (p *Packet) UnmarshalPacket([]byte) error {
	return nil
}
