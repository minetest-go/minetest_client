package commands

import (
	"encoding/binary"
	"fmt"
	"minetest_client/types"
)

type ClientGotBlocks struct {
	Blocks []types.BlockPos
}

func NewClientGotBlocks() *ClientGotBlocks {
	return &ClientGotBlocks{
		Blocks: make([]types.BlockPos, 0),
	}
}

func (p *ClientGotBlocks) AddBlock(pos types.BlockPos) {
	p.Blocks = append(p.Blocks, pos)
}

func (p *ClientGotBlocks) AddBlockPos(x, y, z int16) {
	p.AddBlock(types.BlockPos{
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
		binary.BigEndian.PutUint16(buf[offset:], uint16(bp.PosY))
		offset += 2
		binary.BigEndian.PutUint16(buf[offset:], uint16(bp.PosY))
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
