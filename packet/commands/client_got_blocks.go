package commands

import (
	"encoding/binary"
	"fmt"
)

type BlockPos struct {
	PosX int16
	PosY int16
	PosZ int16
}

func (bp BlockPos) String() string {
	return fmt.Sprintf("{Blockpos %d/%d/%d}", bp.PosX, bp.PosY, bp.PosZ)
}

type ClientGotBlocks struct {
	Blocks []BlockPos
}

func NewClientGotBlocks() *ClientGotBlocks {
	return &ClientGotBlocks{
		Blocks: make([]BlockPos, 0),
	}
}

func (p *ClientGotBlocks) AddBlock(pos BlockPos) {
	p.Blocks = append(p.Blocks, pos)
}

func (p *ClientGotBlocks) AddBlockPos(x, y, z int16) {
	p.AddBlock(BlockPos{
		PosX: x,
		PosY: y,
		PosZ: z,
	})
}

func (p *ClientGotBlocks) GetCommandId() uint16 {
	return ClientCommandGotBlocks
}

func (p *ClientGotBlocks) MarshalPacket() ([]byte, error) {
	buf := make([]byte, 1+(len(p.Blocks)*6))
	buf[0] = uint8(len(p.Blocks))
	offset := 1
	for _, bp := range p.Blocks {
		binary.BigEndian.PutUint16(buf[offset:], uint16(bp.PosX))
		offset += 2
		binary.BigEndian.PutUint16(buf[offset:], uint16(bp.PosX))
		offset += 2
		binary.BigEndian.PutUint16(buf[offset:], uint16(bp.PosX))
		offset += 2
	}

	return buf, nil
}

func (p *ClientGotBlocks) UnmarshalPacket([]byte) error {
	return nil
}

func (p *ClientGotBlocks) String() string {
	return fmt.Sprintf("{ClientGotBlocks blocks=%s}", p.Blocks)
}
