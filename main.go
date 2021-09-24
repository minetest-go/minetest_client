package main

import (
	"bufio"
	"fmt"
	"minetest_client/packet"
	"minetest_client/packet/commands"
	"net"
)

func main() {
	fmt.Println("ok")

	peerInit := commands.NewClientPeerInit()
	p := packet.CreatePacket(packet.Reliable, packet.Original, 0, 65500, peerInit)
	data, err := p.MarshalPacket()
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	conn, err := net.Dial("udp", "127.0.0.1:30000")
	if err != nil {
		panic(err)
	}

	_, err = conn.Write(data)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 4096)
	len, err := bufio.NewReader(conn).Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println(buf[:len])

}
