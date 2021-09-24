package main

import (
	"minetest_client/packet"
	"minetest_client/packet/commands"
	"time"
)

func main() {
	client := NewClient("127.0.0.1", 30000)
	err := client.Start()
	if err != nil {
		panic(err)
	}

	err = client.Send(packet.CreateReliable(0, 65500, commands.NewClientPeerInit()))
	if err != nil {
		panic(err)
	}

	err = client.Send(packet.CreateOriginal(0, 65500, commands.NewClientInit("test")))
	if err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)
}
