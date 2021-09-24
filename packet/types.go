package packet

type Serializeable interface {
	MarshalPacket() ([]byte, error)
	UnmarshalPacket([]byte) error
}

type Command interface {
	Serializeable
	GetCommandId() uint8
}
