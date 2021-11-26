package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	var host, username, password string
	var port int
	var stalk, downloadmedia, help bool

	flag.StringVar(&host, "host", "127.0.0.1", "The hostname")
	flag.IntVar(&port, "port", 30000, "The portname")
	flag.BoolVar(&help, "help", false, "Shows the help")
	flag.StringVar(&username, "username", "test", "The username")
	flag.StringVar(&password, "password", "enter", "The password")
	flag.BoolVar(&stalk, "stalk", false, "Stalk mode: don't really join, just listen")
	flag.BoolVar(&downloadmedia, "media", false, "Download media")
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	fmt.Printf("Connecting to '%s:%d' with username '%s'\n", host, port, username)

	client := NewClient(host, port)

	ch := &ClientHandler{
		Client:        client,
		Username:      username,
		Password:      password,
		StalkMode:     stalk,
		DownloadMedia: downloadmedia,
	}
	client.SetServerCommandHandler(ch)

	err := client.Start()
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)
	err = client.Init()
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("Sending disconnect")
	err = client.Stop()
	if err != nil {
		panic(err)
	}
}
