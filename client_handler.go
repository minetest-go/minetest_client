package main

import (
	"fmt"
	"minetest_client/commands"
	"minetest_client/packet"
	"minetest_client/srp"
	"os"
	"time"
)

type ClientHandler struct {
	Client        *Client
	Username      string
	DownloadMedia bool
	Password      string
	StalkMode     bool
	SRPPubA       []byte
	SRPPrivA      []byte
	MediaHashes   map[string][]byte
}

func (ch *ClientHandler) HandlerLoop() {
	cmd_chan := make(chan commands.Command, 500)
	ch.Client.AddListener(cmd_chan)

	for o := range cmd_chan {
		err := ch.handleCommand(o)
		if err != nil {
			panic(err)
		}
	}
}

func (ch *ClientHandler) handleCommand(o interface{}) error {
	switch cmd := o.(type) {
	case *commands.ServerSetPeer:
		fmt.Printf("Received set_peerid: %d\n", cmd.PeerID)
		time.Sleep(1 * time.Second)
		err := ch.Client.SendOriginalCommand(commands.NewClientInit(ch.Username))
		if err != nil {
			return err
		}

	case *commands.ServerHello:
		packet.ResetSeqNr(65500)
		if ch.StalkMode {
			return nil
		}

		if cmd.AuthMechanismSRP {
			// existing client
			var err error
			ch.SRPPubA, ch.SRPPrivA, err = srp.InitiateHandshake()
			if err != nil {
				return err
			}

			fmt.Println("Sending SRP bytes A")
			err = ch.Client.SendCommand(commands.NewClientSRPBytesA(ch.SRPPubA))
			if err != nil {
				return err
			}
		}

		if cmd.AuthMechanismFirstSRP {
			// new client
			salt, verifier, err := srp.NewClient([]byte(ch.Username), []byte(ch.Password))
			if err != nil {
				return err
			}

			fmt.Println("Sending first SRP")
			err = ch.Client.SendCommand(commands.NewClientFirstSRP(salt, verifier))
			if err != nil {
				return err
			}
		}
	case *commands.ServerAccessDenied:
		fmt.Println("Access denied")

	case *commands.ServerSRPBytesSB:
		identifier := []byte(ch.Username)
		passphrase := []byte(ch.Password)

		clientK, err := srp.CompleteHandshake(ch.SRPPubA, ch.SRPPrivA, identifier, passphrase, cmd.BytesS, cmd.BytesB)
		if err != nil {
			return err
		}

		proof := srp.ClientProof(identifier, cmd.BytesS, ch.SRPPubA, cmd.BytesB, clientK)

		fmt.Println("Sending SRP bytes M")
		err = ch.Client.SendCommand(commands.NewClientSRPBytesM(proof))
		if err != nil {
			return err
		}

	case *commands.ServerAuthAccept:
		fmt.Println("Sending INIT2")
		err := ch.Client.SendCommand(commands.NewClientInit2())
		if err != nil {
			return err
		}

	case *commands.ServerAnnounceMedia:
		fmt.Printf("Server announces media: %d assets\n", cmd.FileCount)
		ch.MediaHashes = cmd.Hashes

		if !ch.DownloadMedia {
			return nil
		}

		_, err := os.Stat("media")
		if os.IsNotExist(err) {
			err := os.Mkdir("media", 0755)
			if err != nil {
				return err
			}
		}

		fmt.Printf("Sending REQUEST_MEDIA len=%d\n", len(ch.MediaHashes))
		files := make([]string, 0)
		for name := range ch.MediaHashes {
			//fmt.Printf("Name: '%s'\n", name)

			_, err := os.Stat("media/" + name)
			if os.IsNotExist(err) {
				files = append(files, name)
			}

		}

		if len(files) > 0 {
			reqmedia_cmd := commands.NewClientRequestMedia(files)
			err = ch.Client.SendCommand(reqmedia_cmd)
			if err != nil {
				return err
			}
		}

	case *commands.ServerMedia:
		fmt.Printf("Server media: %s\n", cmd)

		for name, data := range cmd.Files {
			_, err := os.Stat("media/" + name)
			if os.IsNotExist(err) {
				err = os.WriteFile("media/"+name, data, 0644)
				if err != nil {
					return err
				}
			}
		}

	case *commands.ServerCSMRestrictionFlags:
		fmt.Println("Server sends csm restriction flags")

		fmt.Println("Sending CLIENT_READY")
		err := ch.Client.SendCommand(commands.NewClientReady(5, 5, 5, "mt-bot", 4))
		if err != nil {
			return err
		}

		fmt.Println("Sending PLAYERPOS")
		ppos := commands.NewClientPlayerPos()
		err = ch.Client.SendCommand(ppos)
		if err != nil {
			return err
		}

	case *commands.ServerBlockData:
		gotblocks := commands.NewClientGotBlocks()
		gotblocks.AddBlock(cmd.Pos)

		err := ch.Client.SendCommand(gotblocks)
		if err != nil {
			return err
		}
	case *commands.ServerTimeOfDay:
		fmt.Printf("Time of day: %d\n", cmd.TimeOfDay)
	case *commands.ServerChatMessage:
		fmt.Printf("Chat: '%s'\n", cmd.Message)
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
		err := ch.Client.SendOriginalCommand(ppos)
		if err != nil {
			return err
		}
	}

	return nil
}
