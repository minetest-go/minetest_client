package commandclient

import (
	"errors"

	"github.com/minetest-go/minetest_client/commands"
)

func ClientReady(cc *CommandClient) error {
	for o := range cc.CommandChannel() {
		switch o.(type) {
		case *commands.ServerCSMRestrictionFlags:
			err := cc.SendCommand(commands.NewClientReady(5, 5, 5, "mt-bot", 4))
			if err != nil {
				return err
			}

			ppos := commands.NewClientPlayerPos()
			err = cc.SendCommand(ppos)
			return err
		}
	}

	return errors.New("channel closed")
}
