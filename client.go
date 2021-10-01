package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"minetest_client/packet"
	"minetest_client/packet/commands"
	"net"
)

type ClientPacketListener interface {
	OnPacketReceive(c *Client, p *packet.Packet)
}

type ClientCommandListener interface {
	OnCommandReceive(c *Client, cmd packet.Command)
}

type Client struct {
	conn          net.Conn
	Host          string
	Port          int
	PeerID        uint16
	listeners     []ClientPacketListener
	cmd_listeners []ClientCommandListener
	sph           *packet.SplitpacketHandler
}

func NewClient(host string, port int) *Client {
	return &Client{
		Host:          host,
		Port:          port,
		listeners:     make([]ClientPacketListener, 0),
		cmd_listeners: make([]ClientCommandListener, 0),
		sph:           packet.NewSplitPacketHandler(),
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

func (c *Client) AddCommandListener(l ClientCommandListener) {
	c.cmd_listeners = append(c.cmd_listeners, l)
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

	if p.PacketType == packet.Reliable || p.PacketType == packet.Original {
		if p.ControlType == packet.SetPeerID {
			c.PeerID = p.PeerID
			cmd := &commands.ServerSetPeer{
				PeerID: p.PeerID,
			}

			for _, cmd_listener := range c.cmd_listeners {
				cmd_listener.OnCommandReceive(c, cmd)
			}
		}
	}

	for _, listener := range c.listeners {
		listener.OnPacketReceive(c, p)
	}

	if p.Command != nil && (p.SubType == packet.Reliable || p.SubType == packet.Original) {
		for _, cmd_listener := range c.cmd_listeners {
			cmd_listener.OnCommandReceive(c, p.Command)
		}
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

			//fmt.Printf("Received and assembled command: %s\n", cmd)

			if cmd != nil {
				for _, cmd_listener := range c.cmd_listeners {
					cmd_listener.OnCommandReceive(c, cmd)
				}
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
