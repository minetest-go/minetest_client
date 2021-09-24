package main

import (
	"bufio"
	"fmt"
	"minetest_client/packet"
	"net"
)

type ClientEvent int

const (
	InitDone ClientEvent = iota
)

type ClientEventListener interface {
	OnClientEvent(event ClientEvent)
}

type Client struct {
	conn      net.Conn
	Host      string
	Port      int
	PeerID    uint16
	SeqNr     uint16
	listeners []ClientEventListener
}

func NewClient(host string, port int) *Client {
	return &Client{
		Host:      host,
		Port:      port,
		SeqNr:     65500,
		listeners: make([]ClientEventListener, 0),
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

func (c *Client) AddListener(listener ClientEventListener) {
	c.listeners = append(c.listeners, listener)
}

func (c *Client) EmitEvent(event ClientEvent) {
	for _, listener := range c.listeners {
		listener.OnClientEvent(event)
	}
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

	if p.PacketType == packet.Reliable {
		if p.ControlType == packet.SetPeerID {
			fmt.Printf("Got setpeerid command, data: %s", fmt.Sprint(p.Payload))
		}

		// send ack
		err := c.Send(packet.CreateControl(c.PeerID, p.SeqNr, packet.Ack))
		if err != nil {
			panic(err)
		}

		if p.CommandID == 1 {
			c.EmitEvent(InitDone)
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
