package main

import (
	"fmt"
	"minetest_client/packet"
	"time"
)

func main() {
	client := NewClient("127.0.0.1", 30000)
	ch := &ClientHandler{
		Username: "test",
		Password: "enter",
		client:   client,
	}
	client.AddListener(ch)

	err := client.Start()
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)
	err = ch.Init()
	if err != nil {
		panic(err)
	}

	time.Sleep(300 * time.Second)

	fmt.Println("Sending disconnect")
	err = client.Send(packet.CreateControl(ch.peerID, packet.Disco))
	if err != nil {
		panic(err)
	}
}
