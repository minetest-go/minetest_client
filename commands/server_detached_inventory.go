package commands

import (
	"fmt"
)

type ServerDetachedInventory struct {
	Inventory string
}

func (p *ServerDetachedInventory) GetCommandId() uint16 {
	return ServerCommandDetachedInventory
}

func (p *ServerDetachedInventory) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerDetachedInventory) UnmarshalPacket(payload []byte) error {
	p.Inventory = string(payload[6:])
	return nil
}

func (p *ServerDetachedInventory) String() string {
	return fmt.Sprintf("{ServerDetachedInventory Inventory=%s}", p.Inventory)
}
