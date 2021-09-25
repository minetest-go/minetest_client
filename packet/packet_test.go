package packet

import (
	"fmt"
	"minetest_client/packet/commands"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReliableSRPBytesA(t *testing.T) {
	bytes_a := make([]byte, 256)

	p := CreateReliable(2, 65500, commands.NewClientSRPBytesA(bytes_a))
	data, err := p.MarshalPacket()

	assert.NoError(t, err)
	assert.NotNil(t, data)
	fmt.Println(data)

	// packet length
	assert.Equal(t, 272, len(data))

	// protocol id
	assert.Equal(t, uint8(79), data[0])
	assert.Equal(t, uint8(69), data[1])
	assert.Equal(t, uint8(116), data[2])
	assert.Equal(t, uint8(3), data[3])

	// peer id
	assert.Equal(t, uint8(0), data[4])
	assert.Equal(t, uint8(2), data[5])

	// channel
	assert.Equal(t, uint8(1), data[6])

	// packet type
	assert.Equal(t, uint8(Reliable), data[7])

	// sequence number
	assert.Equal(t, uint8(255), data[8])
	assert.Equal(t, uint8(220), data[9])

	// subtype
	assert.Equal(t, uint8(Original), data[10])

	// command id
	assert.Equal(t, uint8(0), data[11])
	assert.Equal(t, uint8(0x51), data[12])

	// payload
	// TODO
}

func TestReliableFirstSRP(t *testing.T) {
	salt := make([]byte, 16)
	verifier := make([]byte, 256)

	p := CreateReliable(2, 65500, commands.NewClientFirstSRP(salt, verifier))
	data, err := p.MarshalPacket()

	assert.NoError(t, err)
	assert.NotNil(t, data)
	fmt.Println(data)

	// packet length
	assert.Equal(t, 290, len(data))

	// protocol id
	assert.Equal(t, uint8(79), data[0])
	assert.Equal(t, uint8(69), data[1])
	assert.Equal(t, uint8(116), data[2])
	assert.Equal(t, uint8(3), data[3])

	// peer id
	assert.Equal(t, uint8(0), data[4])
	assert.Equal(t, uint8(2), data[5])

	// channel
	assert.Equal(t, uint8(1), data[6])

	// packet type
	assert.Equal(t, uint8(Reliable), data[7])

	// sequence number
	assert.Equal(t, uint8(255), data[8])
	assert.Equal(t, uint8(220), data[9])

	// subtype
	assert.Equal(t, uint8(Original), data[10])

	// command id
	assert.Equal(t, uint8(0), data[11])
	assert.Equal(t, uint8(0x50), data[12])

	// payload
	// TODO
}

func TestReliablePeerInit(t *testing.T) {
	p := CreateReliable(2, 65500, commands.NewClientPeerInit())
	data, err := p.MarshalPacket()

	assert.NoError(t, err)
	assert.NotNil(t, data)
	fmt.Println(data)

	// packet length
	assert.Equal(t, 13, len(data))

	// protocol id
	assert.Equal(t, uint8(79), data[0])
	assert.Equal(t, uint8(69), data[1])
	assert.Equal(t, uint8(116), data[2])
	assert.Equal(t, uint8(3), data[3])

	// peer id
	assert.Equal(t, uint8(0), data[4])
	assert.Equal(t, uint8(2), data[5])

	// channel
	assert.Equal(t, uint8(1), data[6])

	// packet type
	assert.Equal(t, uint8(Reliable), data[7])

	// sequence number
	assert.Equal(t, uint8(255), data[8])
	assert.Equal(t, uint8(220), data[9])

	// subtype
	assert.Equal(t, uint8(Original), data[10])

	// payload
	assert.Equal(t, uint8(0), data[11])
	assert.Equal(t, uint8(0), data[12])
}
