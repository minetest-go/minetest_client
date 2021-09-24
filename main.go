package main

import (
	"minetest_client/packet"
	"minetest_client/packet/commands"
	"time"
)

func main() {
	client := NewClient("edgy1.net", 30025)
	err := client.Start()
	if err != nil {
		panic(err)
	}

	err = client.Send(packet.CreateReliable(0, 65500, commands.NewClientPeerInit()))
	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)

	err = client.Send(packet.CreateOriginal(0, 65500, commands.NewClientInit("test")))
	if err != nil {
		panic(err)
	}

	time.Sleep(60 * time.Second)

	err = client.Send(packet.CreateControl(0, 65500, packet.Disco))
	if err != nil {
		panic(err)
	}
}
