package commandclient

import (
	"errors"
	"time"

	"github.com/minetest-go/minetest_client/commands"
)

func Init(cc *CommandClient, username string) error {
	for o := range cc.CommandChannel() {
		switch o.(type) {
		case *commands.ServerSetPeer:
			time.Sleep(1 * time.Second)
			err := cc.SendOriginalCommand(commands.NewClientInit(username))
			return err
		}
	}

	return errors.New("channel closed")
}
