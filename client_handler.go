package main

import (
	"fmt"
	"minetest_client/packet"
	"minetest_client/packet/commands"
	"minetest_client/srp"
	"time"
)

type ClientHandler struct {
	Username string
	Password string
	peerID   uint16
	client   *Client
	SRPPubA  []byte
	SRPPrivA []byte
}

func (ch *ClientHandler) Init() error {
	return ch.client.Send(packet.CreateReliable(0, commands.NewClientPeerInit()))
}

func (ch *ClientHandler) OnPacketReceive(p *packet.Packet) {
	if p.PacketType == packet.Reliable {
		if p.ControlType == packet.SetPeerID {
			ch.peerID = p.PeerID

			go func() {
				// deferred init
				time.Sleep(2 * time.Second)
				fmt.Println("Sending INIT")
				err := ch.client.Send(packet.CreateOriginal(ch.peerID, commands.NewClientInit(ch.Username)))
				if err != nil {
					panic(err)
				}
			}()
		}

		// send ack
		err := ch.client.Send(packet.CreateControlAck(ch.peerID, p))
		if err != nil {
			panic(err)
		}

		if p.SubType == packet.Split {
			// don't process split packets
			return
		}

		if p.CommandID == commands.ServerCommandHello {
			packet.ResetSeqNr(65500)
			hello_cmd, ok := p.Command.(*commands.ServerHello)
			if !ok {
				panic("invalid type")
			}

			if hello_cmd.AuthMechanismSRP {
				// existing client
				var err error
				ch.SRPPubA, ch.SRPPrivA, err = srp.InitiateHandshake()
				if err != nil {
					panic(err)
				}

				fmt.Println("Sending SRP bytes A")
				err = ch.client.Send(packet.CreateReliable(ch.peerID, commands.NewClientSRPBytesA(ch.SRPPubA)))
				if err != nil {
					panic(err)
				}
			}

			if hello_cmd.AuthMechanismFirstSRP {
				// new client
				salt, verifier, err := srp.NewClient([]byte(ch.Username), []byte(ch.Password))
				if err != nil {
					panic(err)
				}

				fmt.Println("Sending first SRP")
				err = ch.client.Send(packet.CreateReliable(ch.peerID, commands.NewClientFirstSRP(salt, verifier)))
				if err != nil {
					panic(err)
				}
			}
		}

		if p.CommandID == commands.ServerCommandSRPBytesSB {
			sb_cmd, ok := p.Command.(*commands.ServerSRPBytesSB)
			if !ok {
				panic("invalid type")
			}

			identifier := []byte(ch.Username)
			passphrase := []byte(ch.Password)

			clientK, err := srp.CompleteHandshake(ch.SRPPubA, ch.SRPPrivA, identifier, passphrase, sb_cmd.BytesS, sb_cmd.BytesB)
			if err != nil {
				panic(err)
			}

			proof := srp.ClientProof(identifier, sb_cmd.BytesS, ch.SRPPubA, sb_cmd.BytesB, clientK)

			fmt.Println("Sending SRP bytes M")
			err = ch.client.Send(packet.CreateReliable(ch.peerID, commands.NewClientSRPBytesM(proof)))
			if err != nil {
				panic(err)
			}
		}

		if p.CommandID == commands.ServerCommandAuthAccept {
			fmt.Println("Sending INIT2")
			err = ch.client.Send(packet.CreateReliable(ch.peerID, commands.NewClientInit2()))
			if err != nil {
				panic(err)
			}
		}

		if p.CommandID == commands.ServerCommandAnnounceMedia {
			fmt.Println("Server announces media")
		}

		if p.CommandID == commands.ServerCommandCSMRestrictionFlags {
			fmt.Println("Server sends csm restriction flags")

			fmt.Println("Sending CLIENT_READY")
			err = ch.client.Send(packet.CreateReliable(ch.peerID, commands.NewClientReady(5, 5, 5, "mt-bot", 4)))
			if err != nil {
				panic(err)
			}

			fmt.Println("Sending PLAYERPOS")
			ppos := commands.NewClientPlayerPos()
			err = ch.client.Send(packet.CreateReliable(ch.peerID, ppos))
			if err != nil {
				panic(err)
			}

		}

		if p.CommandID == commands.ServerCommandItemDefinitions {
			fmt.Println("Server sends item definitions")
		}

		if p.CommandID == commands.ServerCommandNodeDefinitions {
			fmt.Println("Server sends node definitions")
		}
	}
}
