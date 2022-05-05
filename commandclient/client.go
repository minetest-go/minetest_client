package commandclient

import (
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/minetest-go/minetest_client/commands"
	"github.com/minetest-go/minetest_client/packet"
)

type CommandClient struct {
	conn          net.Conn
	Host          string
	Port          int
	PeerID        uint16
	sph           *packet.SplitpacketHandler
	netrx         chan []byte
	listeners     []chan commands.Command
	listener_lock *sync.RWMutex
}

func NewCommandClient(host string, port int) *CommandClient {
	return &CommandClient{
		Host:          host,
		Port:          port,
		sph:           packet.NewSplitPacketHandler(),
		netrx:         make(chan []byte, 1000),
		listeners:     make([]chan commands.Command, 0),
		listener_lock: &sync.RWMutex{},
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
	return c.conn.Close()
}

func (c *CommandClient) AddListener(ch chan commands.Command) {
	c.listener_lock.Lock()
	defer c.listener_lock.Unlock()
	c.listeners = append(c.listeners, ch)
}

func (c *CommandClient) RemoveListener(ch chan commands.Command) {
	c.listener_lock.Lock()
	defer c.listener_lock.Unlock()
	newlisteners := make([]chan commands.Command, 0)
	for _, l := range c.listeners {
		if l != ch {
			newlisteners = append(newlisteners, l)
		}
	}
	c.listeners = newlisteners
}

func (c *CommandClient) emitCommand(cmd commands.Command) {
	c.listener_lock.RLock()
	defer c.listener_lock.RUnlock()

	for _, ch := range c.listeners {
		select {
		case ch <- cmd:
		default:
		}
	}
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
	cmd, err := commands.Parse(payload)
	if err != nil {
		return err
	}
	c.emitCommand(cmd)
	return nil
}

func (c *CommandClient) onReceive(p *packet.Packet) error {
	//fmt.Printf("Received packet: %s\n", p)

	if p.PacketType == packet.Reliable || p.PacketType == packet.Original {
		if p.ControlType == packet.SetPeerID {
			c.PeerID = p.PeerID
			cmd := &commands.ServerSetPeer{
				PeerID: p.PeerID,
			}

			c.emitCommand(cmd)
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
