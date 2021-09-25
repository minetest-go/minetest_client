package main

import (
	"minetest_client/packet"
	"time"
)

func main() {
	//client := NewClient("pandorabox.io", 30000)
	client := NewClient("127.0.0.1", 30000)
	ch := &ClientHandler{
		Username: "test2",
		Password: "enter",
		client:   client,
	}
	client.AddListener(ch)

	err := client.Start()
	if err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)

	err = client.Send(packet.CreateControl(ch.peerID, 0, packet.Disco))
	if err != nil {
		panic(err)
	}
}
