package packet

type Packet interface {
	MarshalPacket() ([]byte, error)
	UnmarshalPacket([]byte) error
}
