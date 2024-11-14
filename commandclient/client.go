package commandclient

import (
	"errors"
	"fmt"
	"net"

	"github.com/minetest-go/minetest_client/commands"
	"github.com/minetest-go/minetest_client/packet"
)

var ErrTimeout = errors.New("timeout")

type CommandClient struct {
	conn     net.Conn
	Host     string
	Port     int
	PeerID   uint16
	sph      *packet.SplitpacketHandler
	netrx    chan []byte
	cmd_chan chan commands.Command
}

func NewCommandClient(host string, port int) *CommandClient {
	return &CommandClient{
		Host:     host,
		Port:     port,
		sph:      packet.NewSplitPacketHandler(),
		netrx:    make(chan []byte, 1000),
		cmd_chan: make(chan commands.Command, 1000),
	}
}

func (c *CommandClient) Connect() error {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		return err
	}
	c.conn = conn
	go c.rxLoop()
	go c.parseLoop()

	peerInit := packet.CreateReliable(0, []byte{0, 0})
	peerInit.Channel = 0
	return c.Send(peerInit)
}

func (c *CommandClient) Disconnect() error {
	err := c.Send(packet.CreateControl(c.PeerID, packet.Disco))
	if err != nil {
		return err
	}
	close(c.netrx)
	close(c.cmd_chan)
	return c.conn.Close()
}

func (c *CommandClient) CommandChannel() chan commands.Command {
	return c.cmd_chan
}

func (c *CommandClient) SendOriginalCommand(cmd commands.Command) error {
	//fmt.Printf("Sending original command: %s\n", cmd)

	payload, err := commands.CreatePayload(cmd)
	if err != nil {
		return err
	}

	pkg := packet.CreateOriginal(c.PeerID, payload)
	return c.Send(pkg)
}

func (c *CommandClient) SendCommand(cmd commands.Command) error {
	//fmt.Printf("Sending command: %s\n", cmd)

	payload, err := commands.CreatePayload(cmd)
	if err != nil {
		return err
	}

	if len(payload) < packet.MaxPacketLength {
		// one packet
		pkg := packet.CreateReliable(c.PeerID, payload)
		return c.Send(pkg)

	} else {
		// split packet
		pkgs, err := c.sph.SplitPayload(payload)
		if err != nil {
			return err
		}

		for _, pkg := range pkgs {
			pkg.PeerID = c.PeerID
			pkg.Channel = 1
			pkg.SeqNr = packet.NextSequenceNr()

			err = c.Send(pkg)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func (c *CommandClient) Send(packet *packet.Packet) error {
	data, err := packet.MarshalPacket()
	if err != nil {
		return err
	}
	//fmt.Printf("Sending packet: %s\n", packet)
	//fmt.Printf("Sending raw: %s\n", fmt.Sprint(data))

	_, err = c.conn.Write(data)
	return err
}

func (c *CommandClient) handleCommandPayload(payload []byte) error {
	//fmt.Printf("Received bytes: len=%d, cmdId=%d\n", len(payload), commands.GetCommandID(payload))

	cmd, err := commands.Parse(payload)
	if err != nil {
		return err
	}

	c.cmd_chan <- cmd
	return nil
}

func (c *CommandClient) onReceive(p *packet.Packet) error {
	//fmt.Printf("Packet: %s\n", p.String())
	if p.PacketType == packet.Reliable || p.PacketType == packet.Original {
		if p.ControlType == packet.SetPeerID {
			c.PeerID = p.PeerID
			cmd := &commands.ServerSetPeer{PeerID: p.PeerID}

			// send as raw payload to potential listeners
			payload, err := commands.CreatePayload(cmd)
			if err != nil {
				return fmt.Errorf("peerId marshal error: %v", err)
			}

			err = c.handleCommandPayload(payload)
			if err != nil {
				return fmt.Errorf("handleCommandPayload error: %v", err)
			}

			c.cmd_chan <- cmd
		}
	}

	// send ack
	if p.PacketType == packet.Reliable {
		ack := packet.CreateControlAck(c.PeerID, p)
		ack.Channel = p.Channel
		if err := c.Send(ack); err != nil {
			return err
		}
	}

	if p.SubType == packet.Reliable || p.SubType == packet.Original {
		if err := c.handleCommandPayload(p.Payload); err != nil {
			return err
		}
	}

	if p.SubType == packet.Split {
		//shove into list
		data := c.sph.AddPacket(p.SplitPayload)

		if data != nil {
			if err := c.handleCommandPayload(data); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *CommandClient) parseLoop() {
	for buf := range c.netrx {
		//fmt.Printf("Received raw: %s\n", fmt.Sprint(buf))

		p, err := packet.Parse(buf)
		if err != nil {
			panic(err)
		}

		err = c.onReceive(p)
		if err != nil {
			panic(err)
		}
	}
}

func (c *CommandClient) rxLoop() {
	for {
		buf := make([]byte, 1024)
		len, err := c.conn.Read(buf)
		if errors.Is(err, net.ErrClosed) {
			return
		}

		if err != nil {
			panic(err)
		}

		c.netrx <- buf[:len]
	}
}
