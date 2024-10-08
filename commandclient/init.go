package commandclient

import (
	"errors"
	"time"

	"github.com/minetest-go/minetest_client/commands"
)

func Init(cc *CommandClient, username string) error {
	ch := make(chan commands.Command, 100)
	cc.AddListener(ch)
	defer cc.RemoveListener(ch)

	for o := range ch {
		switch o.(type) {
		case *commands.ServerSetPeer:
			time.Sleep(1 * time.Second)
			err := cc.SendOriginalCommand(commands.NewClientInit(username))
			return err
		}
	}

	return errors.New("channel closed")
}
