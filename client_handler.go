package main

import (
	"fmt"
	"minetest_client/commands"
	"minetest_client/packet"
	"minetest_client/srp"
	"time"
)

type ClientHandler struct {
	Username  string
	Password  string
	StalkMode bool
	SRPPubA   []byte
	SRPPrivA  []byte
}

func (ch *ClientHandler) Init(c *Client) error {
	peerInit := packet.CreateReliable(0, commands.NewClientPeerInit())
	peerInit.Channel = 0
	return c.Send(peerInit)
}

func (ch *ClientHandler) OnCommandReceive(c *Client, cmd packet.Command) {
	switch cmd.GetCommandId() {
	case commands.ServerCommandSetPeer:
		go func() {
			time.Sleep(1 * time.Second)
			err := c.Send(packet.CreateOriginal(c.PeerID, commands.NewClientInit(ch.Username)))
			if err != nil {
				panic(err)
			}
		}()
	case commands.ServerCommandHello:
		if ch.StalkMode {
			return
		}
		packet.ResetSeqNr(65500)
		hello_cmd, ok := cmd.(*commands.ServerHello)
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
			err = c.Send(packet.CreateReliable(c.PeerID, commands.NewClientSRPBytesA(ch.SRPPubA)))
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
			err = c.Send(packet.CreateReliable(c.PeerID, commands.NewClientFirstSRP(salt, verifier)))
			if err != nil {
				panic(err)
			}
		}
	case commands.ServerCommandSRPBytesSB:
		sb_cmd, ok := cmd.(*commands.ServerSRPBytesSB)
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
		err = c.Send(packet.CreateReliable(c.PeerID, commands.NewClientSRPBytesM(proof)))
		if err != nil {
			panic(err)
		}
	case commands.ServerCommandAuthAccept:
		fmt.Println("Sending INIT2")
		err := c.Send(packet.CreateReliable(c.PeerID, commands.NewClientInit2()))
		if err != nil {
			panic(err)
		}
	case commands.ServerCommandAnnounceMedia:
		fmt.Println("Server announces media")
	case commands.ServerCommandCSMRestrictionFlags:
		fmt.Println("Server sends csm restriction flags")

		fmt.Println("Sending CLIENT_READY")
		err := c.Send(packet.CreateReliable(c.PeerID, commands.NewClientReady(5, 5, 5, "mt-bot", 4)))
		if err != nil {
			panic(err)
		}

		fmt.Println("Sending PLAYERPOS")
		ppos := commands.NewClientPlayerPos()
		err = c.Send(packet.CreateReliable(c.PeerID, ppos))
		if err != nil {
			panic(err)
		}
	case commands.ServerCommandAccessDenied:
		fmt.Println("Server sends ACCESS_DENIED")
	case commands.ServerCommandItemDefinitions:
		fmt.Println("Server sends item definitions")
	case commands.ServerCommandNodeDefinitions:
		fmt.Println("Server sends node definitions")
	case commands.ServerCommandBlockData:
		block_pkg, ok := cmd.(*commands.ServerBlockData)
		if !ok {
			panic("Invalid type")
		}

		//fmt.Printf("Block: '%s'\n", block_pkg)

		gotblocks := commands.NewClientGotBlocks()
		gotblocks.AddBlock(block_pkg.Pos)

		err := c.Send(packet.CreateReliable(c.PeerID, gotblocks))
		if err != nil {
			panic(err)
		}

	case commands.ServerCommandTimeOfDay:
		tod_pkg, ok := cmd.(*commands.ServerTimeOfDay)
		if !ok {
			panic("Invalid type")
		}

		fmt.Printf("Time of day: %d\n", tod_pkg.TimeOfDay)

	case commands.ServerCommandChatMessage:
		chat_pkg, ok := cmd.(*commands.ServerChatMessage)
		if !ok {
			panic("Invalid type")
		}

		fmt.Printf("Chat: '%s'\n", chat_pkg.Message)
	case commands.ServerCommandActiveObjectMessage:
	case commands.ServerCommandAddParticleSpawner:
	}
}
