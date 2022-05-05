package commands

import (
	"encoding/binary"
	"fmt"

	"github.com/minetest-go/minetest_client/types"
)

type ServerBlockData struct {
	Pos types.BlockPos
}

func (p *ServerBlockData) GetCommandId() uint16 {
	return ServerCommandBlockData
}

func (p *ServerBlockData) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerBlockData) UnmarshalPacket(payload []byte) error {
	blockpos := types.BlockPos{}
	blockpos.PosX = int16(binary.BigEndian.Uint16(payload[0:]))
	blockpos.PosY = int16(binary.BigEndian.Uint16(payload[2:]))
	blockpos.PosZ = int16(binary.BigEndian.Uint16(payload[4:]))
	p.Pos = blockpos
	return nil
}

func (p *ServerBlockData) String() string {
	return fmt.Sprintf("{ServerBlockData pos=%s}", p.Pos)
}
