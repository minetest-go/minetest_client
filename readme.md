
Minetest console client, written in go

State: **WIP**

# About

Console client for minetest

Features:
* Log into remote servers
* Ping servers
* Download media
* Listen to various server events

# Usage

Requirements:
* golang >= 1.17

```bash
go build
```

```
# ./minetest_client --help
Usage of ./minetest_client:
  -help
    	Shows the help
  -host string
    	The hostname (default "127.0.0.1")
  -media
    	Download media
  -password string
    	The password (default "enter")
  -ping
    	Just ping the given host:port and exit
  -port int
    	The portname (default 30000)
  -stalk
    	Stalk mode: don't really join, just listen
  -username string
    	The username (default "test")
```

# Api

```golang
package main

import (
	"github.com/minetest-go/minetest_client/commandclient"
)

func main() {
  host := "127.0.0.1"
  port := 30000
  username := "test"
  password := "test"

  client := commandclient.NewCommandClient(host, port)
	err := client.Connect()
	if err != nil {
		panic(err)
	}

	err = commandclient.Init(client, username)
	if err != nil {
		panic(err)
	}

  err = commandclient.Login(client, username, password)
  if err != nil {
    panic(err)
  }

  err = commandclient.ClientReady(client)
  if err != nil {
    panic(err)
  }
}
```

# License

MIT