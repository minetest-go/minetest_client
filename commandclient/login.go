package commandclient

import (
	"errors"

	"github.com/minetest-go/minetest_client/commands"
	"github.com/minetest-go/minetest_client/packet"
	"github.com/minetest-go/minetest_client/srp"
)

var ErrAccessDenied = errors.New("access denied")
var ErrNotRegistered = errors.New("username not registered")

func Login(cc *CommandClient, username, password string, enable_registration bool) error {
	ch := make(chan commands.Command, 100)
	cc.AddListener(ch)
	defer cc.RemoveListener(ch)

	var srppub, srppriv []byte

	for o := range ch {
		switch cmd := o.(type) {
		case *commands.ServerHello:
			packet.ResetSeqNr(65500)

			if cmd.AuthMechanismSRP {
				// existing client
				var err error
				srppub, srppriv, err = srp.InitiateHandshake()
				if err != nil {
					return err
				}

				err = cc.SendCommand(commands.NewClientSRPBytesA(srppub))
				if err != nil {
					return err
				}
			}

			if cmd.AuthMechanismFirstSRP {
				if !enable_registration {
					// registration not enabled, fail
					return ErrNotRegistered
				}

				// new client
				salt, verifier, err := srp.NewClient([]byte(username), []byte(password))
				if err != nil {
					return err
				}

				err = cc.SendCommand(commands.NewClientFirstSRP(salt, verifier))
				if err != nil {
					return err
				}
			}
		case *commands.ServerAccessDenied:
			return ErrAccessDenied

		case *commands.ServerSRPBytesSB:
			identifier := []byte(username)
			passphrase := []byte(password)

			clientK, err := srp.CompleteHandshake(srppub, srppriv, identifier, passphrase, cmd.BytesS, cmd.BytesB)
			if err != nil {
				return err
			}

			proof := srp.ClientProof(identifier, cmd.BytesS, srppub, cmd.BytesB, clientK)

			err = cc.SendCommand(commands.NewClientSRPBytesM(proof))
			if err != nil {
				return err
			}

		case *commands.ServerAuthAccept:
			err := cc.SendCommand(commands.NewClientInit2())
			return err
		}
	}

	return nil
}
