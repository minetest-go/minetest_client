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
	conn            net.Conn
	Host            string
	Port            int
	listeners       []ClientPacketListener
	splitpart_map   map[uint16]*packet.SplitPayload
	splitpart_count uint16
}

func NewClient(host string, port int) *Client {
	return &Client{
		Host:          host,
		Port:          port,
		listeners:     make([]ClientPacketListener, 0),
		splitpart_map: make(map[uint16]*packet.SplitPayload),
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
		if c.splitpart_map[p.SplitPayload.ChunkNumber] == nil {
			c.splitpart_map[p.SplitPayload.ChunkNumber] = p.SplitPayload
			c.splitpart_count++
		}

		if p.SplitPayload.ChunkCount == uint16(c.splitpart_count) {
			//last packet
			payload := []byte{}
			for i := uint16(0); i < c.splitpart_count; i++ {
				if c.splitpart_map[i] == nil {
					panic(fmt.Sprintf("encountered nil splitpacket at i=%d, parts=%d, chunks=%d",
						i, c.splitpart_count, p.SplitPayload.ChunkCount))
				}
				payload = append(payload, c.splitpart_map[i].Data...)
			}

			//reset split vars
			c.splitpart_map = make(map[uint16]*packet.SplitPayload)
			c.splitpart_count = 0

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

			fmt.Printf("Received and assembled packet: %s\n", p)
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
