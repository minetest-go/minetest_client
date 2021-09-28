package main

import (
	"flag"
	"fmt"
	"minetest_client/packet"
	"time"
)

func main() {
	var host, username, password string
	var port int
	var help bool

	flag.StringVar(&host, "host", "127.0.0.1", "The hostname")
	flag.IntVar(&port, "port", 30000, "The portname")
	flag.BoolVar(&help, "help", false, "Shows the help")
	flag.StringVar(&username, "username", "test", "The username")
	flag.StringVar(&password, "password", "enter", "The password")
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	client := NewClient(host, port)
	client.AddListener(ClientAckHandler{})

	ch := &ClientHandler{
		Username: username,
		Password: password,
	}
	client.AddCommandListener(ch)

	err := client.Start()
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)
	err = ch.Init(client)
	if err != nil {
		panic(err)
	}

	time.Sleep(300 * time.Second)

	fmt.Println("Sending disconnect")
	err = client.Send(packet.CreateControl(client.PeerID, packet.Disco))
	if err != nil {
		panic(err)
	}
}
