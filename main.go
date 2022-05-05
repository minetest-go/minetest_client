package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/minetest-go/minetest_client/commandclient"
	"github.com/minetest-go/minetest_client/commands"
)

func main() {
	var host, username, password string
	var port int
	var stalk, downloadmedia, help, ping bool

	flag.StringVar(&host, "host", "127.0.0.1", "The hostname")
	flag.IntVar(&port, "port", 30000, "The portname")
	flag.BoolVar(&help, "help", false, "Shows the help")
	flag.StringVar(&username, "username", "test", "The username")
	flag.StringVar(&password, "password", "enter", "The password")
	flag.BoolVar(&stalk, "stalk", false, "Stalk mode: don't really join, just listen")
	flag.BoolVar(&downloadmedia, "media", false, "Download media")
	flag.BoolVar(&ping, "ping", false, "Just ping the given host:port and exit")
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	if ping {
		fmt.Printf("Pinging to '%s:%d'\n", host, port)
		ping_res, err := Ping(host, port)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Ping successful: delay=%d ms, peer_id=%d\n",
			ping_res.Delay.Milliseconds(), ping_res.PeerID)
		return
	}

	fmt.Printf("Connecting to '%s:%d' with username '%s'\n", host, port, username)

	client := commandclient.NewCommandClient(host, port)

	ch := &ClientHandler{
		Client:        client,
		Username:      username,
		Password:      password,
		StalkMode:     stalk,
		DownloadMedia: downloadmedia,
	}

	cmd_chan := make(chan commands.Command, 500)
	client.AddListener(cmd_chan)
	go ch.HandlerLoop(cmd_chan)

	err := client.Connect()
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("Sending disconnect")
	err = client.Disconnect()
	if err != nil {
		panic(err)
	}
}
