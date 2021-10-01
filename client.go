package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"minetest_client/packet"
	"net"
)

type ClientPacketListener interface {
	OnPacketReceive(p *packet.Packet)
}

type Client struct {
	conn      net.Conn
	Host      string
	Port      int
	listeners []ClientPacketListener
	sph       *packet.SplitpacketHandler
}

func NewClient(host string, port int) *Client {
	return &Client{
		Host:      host,
		Port:      port,
		listeners: make([]ClientPacketListener, 0),
		sph:       packet.NewSplitPacketHandler(),
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

func (c *Client) AddListener(listener ClientPacketListener) {
	c.listeners = append(c.listeners, listener)
}

func (c *Client) Send(packet *packet.Packet) error {
	data, err := packet.MarshalPacket()
	if err != nil {
		return err
	}
	//fmt.Printf("Sending packet: %s\n", packet)
	//fmt.Printf("Sending raw: %s\n", fmt.Sprint(data))

	_, err = c.conn.Write(data)
	return err
}

func (c *Client) onReceive(p *packet.Packet) {
	//fmt.Printf("Received packet: %s\n", p)

	for _, listener := range c.listeners {
		listener.OnPacketReceive(p)
	}

	if p.SubType == packet.Split {
		//shove into list
		data := c.sph.AddPacket(p.SplitPayload)

		if data != nil {
			commandId := binary.BigEndian.Uint16(data[0:])
			cmd, err := packet.CreateCommand(commandId, data[2+4:])
			if err != nil {
				panic(err)
			}

			p.CommandID = commandId
			p.Command = cmd
			p.Payload = data
			p.SubType = packet.Original
			p.SplitPayload = nil

			//fmt.Printf("Received and assembled packet: %s\n", p)
			for _, listener := range c.listeners {
				listener.OnPacketReceive(p)
			}
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

		//fmt.Printf("Received raw: %s\n", fmt.Sprint(buf[:len]))

		p, err := packet.Parse(buf[:len])
		if err != nil {
			panic(err)
		}

		c.onReceive(p)
	}
}
