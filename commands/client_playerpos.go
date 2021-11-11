package commands

import (
	"encoding/binary"
	"fmt"
)

type ClientPlayerPos struct {
	PosX             uint32
	PosY             uint32
	PosZ             uint32
	SpeedX           uint32
	SpeedY           uint32
	SpeedZ           uint32
	Pitch            uint32
	Yaw              uint32
	PressedKeys      uint32
	FOV              uint8
	RequestViewRange uint8
}

func NewClientPlayerPos() *ClientPlayerPos {
	return &ClientPlayerPos{}
}

func (p *ClientPlayerPos) GetCommandId() uint16 {
	return ClientCommandPlayerPos
}

func (p *ClientPlayerPos) MarshalPacket() ([]byte, error) {
	data := make([]byte, 38)
	binary.BigEndian.PutUint32(data[0:], p.PosX)
	binary.BigEndian.PutUint32(data[4:], p.PosY)
	binary.BigEndian.PutUint32(data[8:], p.PosZ)
	binary.BigEndian.PutUint32(data[12:], p.SpeedX)
	binary.BigEndian.PutUint32(data[16:], p.SpeedY)
	binary.BigEndian.PutUint32(data[20:], p.SpeedZ)
	binary.BigEndian.PutUint32(data[24:], p.Pitch)
	binary.BigEndian.PutUint32(data[28:], p.Yaw)
	binary.BigEndian.PutUint32(data[32:], p.PressedKeys)
	data[36] = p.FOV
	data[37] = p.RequestViewRange

	return data, nil
}

func (p *ClientPlayerPos) UnmarshalPacket([]byte) error {
	return nil
}

func (p *ClientPlayerPos) String() string {
	return fmt.Sprintf("{ClientPlayerPos pos=%d/%d/%d speed=%d/%d/%d pitch=%d yaw=%d}",
		p.PosX, p.PosY, p.PosZ, p.SpeedX, p.SpeedY, p.SpeedZ, p.Pitch, p.Yaw)
}
