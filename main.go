package main

import (
	"minetest_client/packet"
	"minetest_client/packet/commands"
	"time"
)

func main() {
	//client := NewClient("pandorabox.io", 30000)
	client := NewClient("127.0.0.1", 30000)
	err := client.Start()
	if err != nil {
		panic(err)
	}

	ch := &ClientHandler{
		Username: "test2",
		Password: "enter",
		client:   client,
	}
	client.AddListener(ch)

	err = client.Send(packet.CreateReliable(0, 65500, commands.NewClientPeerInit()))
	if err != nil {
		panic(err)
	}

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
