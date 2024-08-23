package commands

import (
	"encoding/binary"
	"fmt"

	"github.com/minetest-go/types"
)

type ServerBlockData struct {
	Pos types.Pos
}

func (p *ServerBlockData) GetCommandId() uint16 {
	return ServerCommandBlockData
}

func (p *ServerBlockData) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerBlockData) UnmarshalPacket(payload []byte) error {
	blockpos := types.Pos{
		int(int16(binary.BigEndian.Uint16(payload[0:]))),
		int(int16(binary.BigEndian.Uint16(payload[2:]))),
		int(int16(binary.BigEndian.Uint16(payload[4:]))),
	}
	p.Pos = blockpos
	return nil
}

func (p *ServerBlockData) String() string {
	return fmt.Sprintf("{ServerBlockData pos=%s}", &p.Pos)
}
