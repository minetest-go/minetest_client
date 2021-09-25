package main

import (
	"minetest_client/packet"
	"minetest_client/packet/commands"
	"minetest_client/srp"
	"time"
)

type ClientHandler struct {
	peerID   uint16
	client   *Client
	SRPPubA  []byte
	SRPPrivA []byte
}

func (ch *ClientHandler) OnPacketReceive(p *packet.Packet) {
	if p.PacketType == packet.Reliable {
		if p.ControlType == packet.SetPeerID {
			ch.peerID = p.PeerID

			go func() {
				// deferred init
				time.Sleep(2 * time.Second)
				err := ch.client.Send(packet.CreateOriginal(ch.peerID, 0, commands.NewClientInit("test")))
				if err != nil {
					panic(err)
				}
			}()
		}

		// send ack
		err := ch.client.Send(packet.CreateControl(ch.peerID, p.SeqNr, packet.Ack))
		if err != nil {
			panic(err)
		}

		if p.CommandID == commands.ServerCommandHello {
			var err error
			ch.SRPPubA, ch.SRPPrivA, err = srp.InitiateHandshake()
			if err != nil {
				panic(err)
			}

			err = ch.client.Send(packet.CreateReliable(ch.peerID, 0, commands.NewClientSRPBytesA(ch.SRPPubA)))
			if err != nil {
				panic(err)
			}
		}

		if p.CommandID == commands.ServerCommandSRPBytesSB {
			sb_cmd, ok := p.Command.(*commands.ServerSRPBytesSB)
			if !ok {
				panic("invalid type")
			}

			identifier := []byte("test")
			passphrase := []byte("enter")

			clientK, err := srp.CompleteHandshake(ch.SRPPubA, ch.SRPPrivA, identifier, passphrase, sb_cmd.BytesS, sb_cmd.BytesB)
			if err != nil {
				panic(err)
			}

			proof := srp.ClientProof(identifier, sb_cmd.BytesS, ch.SRPPubA, sb_cmd.BytesB, clientK)

			err = ch.client.Send(packet.CreateReliable(ch.peerID, 0, commands.NewClientSRPBytesM(proof)))
			if err != nil {
				panic(err)
			}
		}

	}
}

func main() {
	//client := NewClient("pandorabox.io", 30000)
	client := NewClient("127.0.0.1", 30000)
	err := client.Start()
	if err != nil {
		panic(err)
	}

	ch := &ClientHandler{client: client}
	client.AddListener(ch)

	err = client.Send(packet.CreateReliable(0, 65500, commands.NewClientPeerInit()))
	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)

	/*
		identifier := []byte("test")
		passphrase := []byte("enter")
		salt, verifier, err := srp.NewClient(identifier, passphrase)
		if err != nil {
			panic(err)
		}

		err = client.Send(packet.CreateReliable(client.PeerID, 0, commands.NewClientFirstSRP(salt, verifier)))
		if err != nil {
			panic(err)
		}
	*/

	time.Sleep(10 * time.Second)

	err = client.Send(packet.CreateControl(ch.peerID, 0, packet.Disco))
	if err != nil {
		panic(err)
	}
}
