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

const (
	Ack       = 0
	SetPeerID = 1
	Ping      = 2
	Disco     = 3
)
