package main

import (
	"minetest_client/packet"
)

type ClientAckHandler struct{}

func (ah ClientAckHandler) OnPacketReceive(c *Client, p *packet.Packet) {
	// send ack
	if p.PacketType == packet.Reliable {
		ack := packet.CreateControlAck(c.PeerID, p)
		ack.Channel = p.Channel
		err := c.Send(ack)
		if err != nil {
			panic(err)
		}
	}
}
