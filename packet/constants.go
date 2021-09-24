package packet

var ProtocolID = []byte{0x4f, 0x45, 0x74, 0x03}

type PacketType byte

const (
	Control  PacketType = 0
	Original PacketType = 1
	Split    PacketType = 2
	Reliable PacketType = 3
)

type ControlType byte

func (t PacketType) String() string {
	switch t {
	case 0:
		return "Control"
	case 1:
		return "Original"
	case 2:
		return "Split"
	case 3:
		return "Reliable"
	}

	return "Unknown"
}

const (
	Ack       = 0
	SetPeerID = 1
	Ping      = 2
	Disco     = 3
)

func (c ControlType) String() string {
	switch c {
	case 0:
		return "Ack"
	case 1:
		return "SetPeerID"
	case 2:
		return "Ping"
	case 3:
		return "Disconnect"
	}

	return "Unknown"
}
