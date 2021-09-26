package main

import (
	"bufio"
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
	SeqNr     uint16
	listeners []ClientPacketListener
}

func NewClient(host string, port int) *Client {
	return &Client{
		Host:      host,
		Port:      port,
		SeqNr:     65500,
		listeners: make([]ClientPacketListener, 0),
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
	if packet.SeqNr == 0 {
		packet.SeqNr = c.SeqNr
		c.SeqNr++
	}

	data, err := packet.MarshalPacket()
	if err != nil {
		return err
	}
	fmt.Printf("Sending packet: %s\n", packet)
	//fmt.Printf("Sending raw: %s\n", fmt.Sprint(data))

	_, err = c.conn.Write(data)
	return err
}

func (c *Client) onReceive(p *packet.Packet) {
	fmt.Printf("Received packet: %s\n", p)

	for _, listener := range c.listeners {
		listener.OnPacketReceive(p)
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
