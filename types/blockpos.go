package types

import "fmt"

type BlockPos struct {
	PosX int16
	PosY int16
	PosZ int16
}

func (bp BlockPos) String() string {
	return fmt.Sprintf("{Blockpos %d/%d/%d}", bp.PosX, bp.PosY, bp.PosZ)
}
