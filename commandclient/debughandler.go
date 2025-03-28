package commandclient

import (
	"errors"
	"fmt"
	"time"

	"github.com/minetest-go/minetest_client/commands"
)

func DebugHandler(cc *CommandClient) error {
	for o := range cc.CommandChannel() {
		switch cmd := o.(type) {
		case *commands.ServerBlockData:
			gotblocks := commands.NewClientGotBlocks()
			gotblocks.AddBlock(cmd.Pos)

			fmt.Printf("Got block: %s\n", &cmd.Pos)
			err := cc.SendCommand(gotblocks)
			if err != nil {
				return err
			}
		case *commands.ServerNodeDefinitions:
			fmt.Printf("Nodedefs: %v\n", cmd.Definitions)
		case *commands.ServerTimeOfDay:
			fmt.Printf("Time of day: %d\n", cmd.TimeOfDay)
		case *commands.ServerChatMessage:
			fmt.Printf("Chat: %s\n", cmd)
		case *commands.ServerMovePlayer:
			fmt.Printf("Move player: '%s'\n", cmd)

			time.Sleep(time.Second * 2)
			fmt.Printf("Sending player pos command\n")
			ppos := commands.NewClientPlayerPos()
			ppos.PosX = uint32(cmd.X)
			ppos.PosY = uint32(cmd.Y) + 50
			ppos.PosZ = uint32(cmd.Z) + 50
			ppos.FOV = 149
			ppos.RequestViewRange = 13
			err := cc.SendOriginalCommand(ppos)
			if err != nil {
				return err
			}
		}
	}

	return errors.New("channel closed")
}
