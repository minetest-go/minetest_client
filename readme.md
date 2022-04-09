
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

# License

MIT