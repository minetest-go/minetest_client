package packet

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type Packet struct {
	PacketType   PacketType
	SubType      PacketType
	ControlType  ControlType
	PeerID       uint16
	SeqNr        uint16
	Channel      uint8
	Payload      []byte
	SplitPayload *SplitPayload
}

type SplitPayload struct {
	SeqNr       uint16
	ChunkCount  uint16
	ChunkNumber uint16
	Data        []byte
}

func (spl *SplitPayload) String() string {
	return fmt.Sprintf("{SplitPayload SeqNr=%d, Chunk=%d/%d, #Data=%d}",
		spl.SeqNr, spl.ChunkNumber+1, spl.ChunkCount, len(spl.Data))
}

func Parse(data []byte) (*Packet, error) {
	p := &Packet{}
	err := p.UnmarshalPacket(data)
	return p, err
}

var seqNr uint16 = 65500 - 1

func ResetSeqNr(value uint16) {
	seqNr = value - 1
}

func NextSequenceNr() uint16 {
	if seqNr >= 65535 {
		seqNr = 0
	} else {
		seqNr++
	}

	return seqNr
}

func CreatePayload(cmd Command) ([]byte, error) {
	inner_payload, err := cmd.MarshalPacket()
	if err != nil {
		return nil, err
	}

	payload := make([]byte, len(inner_payload)+2)
	copy(payload[2:], inner_payload)

	binary.BigEndian.PutUint16(payload[0:], cmd.GetCommandId())

	return payload, nil
}

func CreateReliable(peerId uint16, payload []byte) *Packet {
	return &Packet{
		PacketType: Reliable,
		Payload:    payload,
		SubType:    Original,
		PeerID:     peerId,
		SeqNr:      NextSequenceNr(),
		Channel:    1,
	}
}

func CreateOriginal(peerId uint16, payload []byte) *Packet {
	return &Packet{
		PacketType: Original,
		Payload:    payload,
		PeerID:     peerId,
		Channel:    1,
	}
}

func CreateControlAck(peerId uint16, packet *Packet) *Packet {
	return &Packet{
		PacketType:  Control,
		ControlType: Ack,
		PeerID:      peerId,
		SeqNr:       packet.SeqNr,
		Channel:     0,
	}
}

func CreateControl(peerId uint16, controlType ControlType) *Packet {
	return &Packet{
		PacketType:  Control,
		ControlType: controlType,
		PeerID:      peerId,
		SeqNr:       NextSequenceNr(),
	}
}

func (p *Packet) MarshalPacket() ([]byte, error) {
	packet := make([]byte, 4+2+1+1)
	copy(packet[0:], ProtocolID)
	binary.BigEndian.PutUint16(packet[4:], p.PeerID)
	packet[6] = p.Channel
	packet[7] = byte(p.PacketType)

	if p.PacketType == Reliable {
		bytes := make([]byte, 3)
		binary.BigEndian.PutUint16(bytes, p.SeqNr)
		bytes[2] = byte(p.SubType)

		packet = append(packet, bytes...)
		packet = append(packet, p.Payload...)

	} else if p.PacketType == Control {
		bytes := make([]byte, 3)
		bytes[0] = byte(p.ControlType)
		binary.BigEndian.PutUint16(bytes[1:], p.SeqNr)
		packet = append(packet, bytes...)

	} else if p.PacketType == Original {
		packet = append(packet, p.Payload...)

	}

	return packet, nil
}

func (p *Packet) UnmarshalPacket(data []byte) error {
	if len(data) < 5 {
		return errors.New("invalid packet length")
	}

	for i, sig := range ProtocolID {
		if data[i] != sig {
			return errors.New("invalid protocol_id")
		}
	}

	p.PeerID = binary.BigEndian.Uint16(data[4:])
	p.Channel = data[6]
	p.PacketType = PacketType(data[7])

	if p.PacketType == Reliable {
		seqNr := binary.BigEndian.Uint16(data[8:])
		p.SeqNr = seqNr
		p.SubType = PacketType(data[10])

		switch p.SubType {
		case Control:
			p.ControlType = ControlType(data[11])

			if p.ControlType == SetPeerID {
				p.PeerID = binary.BigEndian.Uint16(data[12:])
			}
		case Split:
			spl := &SplitPayload{}
			spl.SeqNr = binary.BigEndian.Uint16(data[11:])
			spl.ChunkCount = binary.BigEndian.Uint16(data[13:])
			spl.ChunkNumber = binary.BigEndian.Uint16(data[15:])
			spl.Data = data[17:]

			p.SplitPayload = spl
		default:
			//fmt.Printf("Unknown packet: %s\n", fmt.Sprint(data))
			p.Payload = data[11:]
		}
	}

	return nil
}

func (p *Packet) String() string {
	return fmt.Sprintf("{Packet Type: %s, PeerID: %d, Channel: %d, SeqNr: %d,"+
		"Subtype: %s, ControlType: %s, #Payload: %d, SplitPayload: %s}",
		p.PacketType, p.PeerID, p.Channel, p.SeqNr,
		p.SubType, p.ControlType, len(p.Payload), p.SplitPayload)
}
