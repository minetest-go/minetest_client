package main

import (
	"minetest_client/packet"
	"minetest_client/packet/commands"
	"minetest_client/srp"
	"time"
)

type ClientHandler struct {
	peerID uint16
	client *Client
}

func (ch *ClientHandler) OnPacketReceive(p *packet.Packet) {
	if p.PacketType == packet.Reliable {
		if p.ControlType == packet.SetPeerID {
			ch.peerID = p.PeerID
		}

		// send ack
		err := ch.client.Send(packet.CreateControl(ch.peerID, p.SeqNr, packet.Ack))
		if err != nil {
			panic(err)
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

	time.Sleep(1 * time.Second)

	err = client.Send(packet.CreateOriginal(ch.peerID, 0, commands.NewClientInit("test")))
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

	pub_a, _, err := srp.InitiateHandshake()
	if err != nil {
		panic(err)
	}

	err = client.Send(packet.CreateReliable(ch.peerID, 0, commands.NewClientSRPBytesA(pub_a)))
	if err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)

	err = client.Send(packet.CreateControl(ch.peerID, 0, packet.Disco))
	if err != nil {
		panic(err)
	}
}
