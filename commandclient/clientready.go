package commandclient

import (
	"errors"
	"fmt"

	"github.com/minetest-go/minetest_client/commands"
)

func ClientReady(cc *CommandClient) error {
	ch := make(chan commands.Command, 100)
	cc.AddListener(ch)
	defer cc.RemoveListener(ch)

	for o := range ch {
		switch o.(type) {
		case *commands.ServerCSMRestrictionFlags:
			fmt.Println("Server sends csm restriction flags")

			fmt.Println("Sending CLIENT_READY")
			err := cc.SendCommand(commands.NewClientReady(5, 5, 5, "mt-bot", 4))
			if err != nil {
				return err
			}

			fmt.Println("Sending PLAYERPOS")
			ppos := commands.NewClientPlayerPos()
			err = cc.SendCommand(ppos)
			return err
		}
	}

	return errors.New("channel closed")
}
