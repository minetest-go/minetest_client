package commandclient

import (
	"errors"
	"fmt"

	"github.com/minetest-go/minetest_client/commands"
	"github.com/minetest-go/minetest_client/packet"
	"github.com/minetest-go/minetest_client/srp"
)

var ErrAccessDenied = errors.New("access denied")

func Login(cc *CommandClient, username, password string) error {
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

				fmt.Printf("Sending SRP bytes A, len=%d\n", len(srppub))
				err = cc.SendCommand(commands.NewClientSRPBytesA(srppub))
				if err != nil {
					return err
				}
			}

			if cmd.AuthMechanismFirstSRP {
				// new client
				salt, verifier, err := srp.NewClient([]byte(username), []byte(password))
				if err != nil {
					return err
				}

				fmt.Printf("Sending first SRP, salt-len=%d, verifier-len=%d\n", len(salt), len(verifier))
				err = cc.SendCommand(commands.NewClientFirstSRP(salt, verifier))
				if err != nil {
					return err
				}
			}
		case *commands.ServerAccessDenied:
			fmt.Println("Access denied")
			return ErrAccessDenied

		case *commands.ServerSRPBytesSB:
			identifier := []byte(username)
			passphrase := []byte(password)

			clientK, err := srp.CompleteHandshake(srppub, srppriv, identifier, passphrase, cmd.BytesS, cmd.BytesB)
			if err != nil {
				return err
			}

			proof := srp.ClientProof(identifier, cmd.BytesS, srppub, cmd.BytesB, clientK)

			fmt.Printf("Sending SRP bytes M, len=%d\n", len(proof))
			err = cc.SendCommand(commands.NewClientSRPBytesM(proof))
			if err != nil {
				return err
			}

		case *commands.ServerAuthAccept:
			fmt.Println("Sending INIT2")
			err := cc.SendCommand(commands.NewClientInit2())
			return err
		}
	}

	return nil
}
