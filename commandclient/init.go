package commandclient

import (
	"errors"
	"time"

	"github.com/minetest-go/minetest_client/commands"
)

type InitOpts struct {
	ClientInit *commands.ClientInit
}

func Init(cc *CommandClient, username string, opts *InitOpts) error {
	if opts == nil {
		opts = &InitOpts{}
	}
	if opts.ClientInit == nil {
		opts.ClientInit = commands.NewClientInit(username)
	}

	for o := range cc.CommandChannel() {
		switch o.(type) {
		case *commands.ServerSetPeer:
			time.Sleep(1 * time.Second)
			err := cc.SendOriginalCommand(opts.ClientInit)
			return err
		}
	}

	return errors.New("channel closed")
}
