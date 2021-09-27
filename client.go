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
	conn       net.Conn
	Host       string
	Port       int
	listeners  []ClientPacketListener
	splitparts []*packet.SplitPayload
}

func NewClient(host string, port int) *Client {
	return &Client{
		Host:       host,
		Port:       port,
		listeners:  make([]ClientPacketListener, 0),
		splitparts: make([]*packet.SplitPayload, 0),
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
		c.splitparts = append(c.splitparts, p.SplitPayload)

		if p.SplitPayload.ChunkCount == p.SplitPayload.ChunkNumber+1 {
			//last packet
			payload := []byte{}
			for _, sp := range c.splitparts {
				payload = append(payload, sp.Data...)
			}
			c.splitparts = make([]*packet.SplitPayload, 0)

			commandId := binary.BigEndian.Uint16(payload[0:])
			cmd, err := packet.CreateCommand(commandId, payload[2+4:])
			if err != nil {
				panic(err)
			}

			p.CommandID = commandId
			p.Command = cmd
			p.Payload = payload
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
