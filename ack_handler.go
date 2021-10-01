package main

import "minetest_client/packet"

type ClientAckHandler struct{}

func (ah ClientAckHandler) OnPacketReceive(c *Client, p *packet.Packet) {
	// send ack
	if p.PacketType == packet.Reliable {
		err := c.Send(packet.CreateControlAck(c.PeerID, p))
		if err != nil {
			panic(err)
		}
	}
}
