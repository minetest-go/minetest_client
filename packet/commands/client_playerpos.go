package commands

import (
	"encoding/binary"
	"fmt"
	"math"
)

type ClientPlayerPos struct {
	PosX             float32
	PosY             float32
	PosZ             float32
	SpeedX           float32
	SpeedY           float32
	SpeedZ           float32
	Pitch            float32
	Yaw              float32
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
	binary.BigEndian.PutUint32(data[0:], math.Float32bits(p.PosX))
	binary.BigEndian.PutUint32(data[4:], math.Float32bits(p.PosY))
	binary.BigEndian.PutUint32(data[8:], math.Float32bits(p.PosZ))
	binary.BigEndian.PutUint32(data[12:], math.Float32bits(p.SpeedX))
	binary.BigEndian.PutUint32(data[16:], math.Float32bits(p.SpeedY))
	binary.BigEndian.PutUint32(data[20:], math.Float32bits(p.SpeedZ))
	binary.BigEndian.PutUint32(data[24:], math.Float32bits(p.Pitch))
	binary.BigEndian.PutUint32(data[28:], math.Float32bits(p.Yaw))
	binary.BigEndian.PutUint32(data[32:], p.PressedKeys)
	data[36] = p.FOV
	data[37] = p.RequestViewRange

	return data, nil
}

func (p *ClientPlayerPos) UnmarshalPacket([]byte) error {
	return nil
}

func (p *ClientPlayerPos) String() string {
	return fmt.Sprintf("{ClientPlayerPos pos=%f/%f/%f speed=%f/%f/%f pitch=%f yaw=%f}",
		p.PosX, p.PosY, p.PosZ, p.SpeedX, p.SpeedY, p.SpeedZ, p.Pitch, p.Yaw)
}
