package main

import (
	"flag"
	"fmt"
	"minetest_client/packet"
	"time"
)

func main() {
	host := flag.String("host", "127.0.0.1", "The hostname")
	port := flag.Int("port", 30000, "The portname")
	help := flag.Bool("help", false, "Shows the help")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	client := NewClient(*host, *port)
	ch := &ClientHandler{
		Username: "totallynotabot",
		Password: "Endor",
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
