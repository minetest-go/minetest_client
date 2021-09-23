package main

import (
	"fmt"
	"minetest_client/packet"
)

func main() {
	fmt.Println("ok")

	initPacket := packet.NewClientInit("test")
	data, err := initPacket.MarshalPacket()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
