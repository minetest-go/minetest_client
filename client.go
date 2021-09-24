package main

import (
	"bufio"
	"fmt"
	"minetest_client/packet"
	"minetest_client/packet/commands"
	"net"
)

type Client struct {
	conn   net.Conn
	Host   string
	Port   int
	PeerID uint16
}

func NewClient(host string, port int) *Client {
	return &Client{
		Host: host,
		Port: port,
	}
}

func (c *Client) Start() error {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		return err
	}
	c.conn = conn
	go c.rxLoop()
	return nil
}

func (c *Client) Send(packet *packet.Packet) error {
	data, err := packet.MarshalPacket()
	if err != nil {
		return err
	}
	fmt.Printf("Sending packet: %s\n", packet)
	fmt.Printf("Sending raw: %s\n", fmt.Sprint(data))

	_, err = c.conn.Write(data)
	return err
}

func (c *Client) onReceive(p *packet.Packet) {
	fmt.Printf("Received packet: %s\n", p)

	if p.PacketType == packet.Reliable {
		if p.CommandID == 1 {
			// Set peer id
			setpeer, ok := p.Command.(*commands.ServerSetPeer)
			if ok {
				fmt.Printf("Setting PeerID to %d\n", setpeer.PeerID)
				c.PeerID = setpeer.PeerID
			}
		}

		// send ack
		err := c.Send(packet.CreateControl(c.PeerID, p.SeqNr, packet.Ack))
		if err != nil {
			panic(err)
		}
	}
}

func (c *Client) rxLoop() {
	for {
		buf := make([]byte, 4096)
		len, err := bufio.NewReader(c.conn).Read(buf)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Received raw: %s\n", fmt.Sprint(buf[:len]))

		p, err := packet.Parse(buf[:len])
		if err != nil {
			panic(err)
		}

		c.onReceive(p)
	}
}
