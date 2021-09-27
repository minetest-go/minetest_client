package packet

import (
	"fmt"
	"minetest_client/packet/commands"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSplit(t *testing.T) {
	raw_split_pkg := []byte("\x4f\x45\x74\x03\x00\x01" +
		"\x00\x03\xff\xdf\x02\xff\xdc\x00\x20\x00\x00\x00\x3d\x00\x00\x3c" +
		"\x7d\x78\x9c\xcd\x7d\x0b\x90\x24\x49\x79\x5e\xcd\xf4\x4c\xef\xfb" +
		"\x76\xf7\x66\x6f\x77\x67\x9f\x7d\x77\xa0\x63\x79\xd8\x77\xbb\xb7" +
		"\x77\x30\xb6\xb9\x65\xf7\x1e\x9c\xb8\xe5\x36\x6e\x16\x2e\x24\x22" +
		"\xd4\xaa\xe9\xae\x99\x69\x4d\x4f\x57\xbb\xba\x67\xe7\xc6\x8a\x70" +
		"\x40\x84\x2d\xa3\x07\x61\xc9\x52\x38\x02\x3f\x70\xd8\x61\x49\x21" +
		"\x09\x09\x89\x87\x84\x64\x24\x9d\x2d\x84\x25\x38\x40\x3c\xc4\x4b" +
		"\x08\x0c\xb6\x91\xb1\x8d\x15\x48\x21\xc2\xb2\xb1\xe4\xcc\x3f\xb3" +
		"\xaa\xfe\xcc\xfc\x33\xab\xba\xb3\x6a\x58\x22\xd8\x9b\xce\xcc\xfa" +
		"\xff\xff\xfb\xf3\xf5\xe7\x9f\x7f\x66\x06\xb3\x77\x04\x9f\x6b\x06" +
		"\xf2\x7f\x87\xb6\x7b\x51\xbf\xbb\x1e\x0e\xba\x7f\x63\x38\x58\x7b" +
		"\xe4\xcd\x41\xc0\xff\x7f\xa5\xc5\xb2\x3a\xec\xff\x6f\x99\x7f\x64" +
		"\x75\x75\x15\x4a\x36\x82\x3d\x9d\x64\x6b\x73\xa5\xbf\xc3\x7e\xcc" +
		"\xb0\xff\xcf\x06\x8d\x47\x2e\x5d\xba\x14\xcc\x5e\xb9\xc2\x7e\x9d" +
		"\x88\xbb\xdd\xfe\x4e\x7b\x25\x89\xc2\x8d\x70\xa5\x1f\xb5\x57\x76" +
		"\xda\x9c\xac\x2c\xdc\x90\x85\x67\xae\x7c\x3f\xff\x94\x7f\x11\x34" +
		"\x47\x83\x70\x38\x4c\xc9\xcd\x04\x8d\x57\xbf\xf0\xc2\x47\xe1\xaf" +
		"\xe6\x6a\x3f\x1a\xad\xef\x40\x3a\xfc\x2f\x15\xcc\xfc\xf5\xdb\xf0" +
		"\xdf\xbf\x66\xff\x83\x9f\x7f\xbf\xc9\xe8\x84\xbd\x24\x68\xbc\x86" +
		"\xfd\xb3\x87\xfd\xc5\x71\x65\x7f\xa4\x5f\xc2\xd7\x1d\xf8\x62\x26" +
		"\x58\x1c\xc4\xe3\x76\x6f\xd0\xee\x30\xd9\xc7\xbd\x5b\x11\xfb\xfb" +
		"\x56\x34\x18\xc7\x09\x17\x00\x88\x95\xe7\xff\x5f\x19\xff\xc3\x2b" +
		"\x51\x77\xb4\xc4\xfe\x69\xaf\xc4\xe3\x71\xbc\x19\xdc\x75\xfa\x25" +
		"\x37\xaf\xf0\xc4\x0b\xcb\xbd\xcd\x61\x3f\x6a\x5d\x8d\xba\xa7\x1f" +
		"\x0b\x0e\xf2\xa4\x36\xfb\x07\x44\x54\x7e\x29\x72\x0a\x2d\xcc\xd9" +
		"\x55\x3c\x1b\xec\x5b\xed\x87\x9b\x9b\x3c\x91\x6b\x9a\x11\xe1\x4a" +
		"\xec\xac\xc7\x5c\xbd\xb3\xa6\x40\x27\xba\xd1\x6a\xb8\xd5\x1f\xb7" +
		"\x87\xfd\xb0\x13\xb5\x07\x71\x37\x62\x94\x92\x6e\x69\x98\xef\x66" +
		"\x30\x0f\x66\x54\xc7\xf1\x50\xfd\x02\xa9\x77\x42\xb1\x67\x91\xd8" +
		"\x2a\x03\x6f\x99\xbf\xcd\x64\xbe\x0b\x48\xae\x86\x83\xce\x4e\x1b" +
		"\xe9\xe3\x58\x56\x41\x8f\xf3\x2c\x59\x3f\x77\xa6\x35\xd2\x86\x0f" +
		"\xa0\x96\x88\xb4\x6a\xeb\xca\x22\xa1\x37")

	pkg, err := Parse(raw_split_pkg)
	assert.NoError(t, err)
	assert.NotNil(t, pkg)
	assert.NotNil(t, pkg.SplitPayload)
	assert.NotNil(t, pkg.SplitPayload.Data)

	fmt.Println(pkg)
}

func TestReliableSRPBytesA(t *testing.T) {
	bytes_a := make([]byte, 256)

	p := CreateReliable(2, commands.NewClientSRPBytesA(bytes_a))
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

	p := CreateReliable(2, commands.NewClientFirstSRP(salt, verifier))
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

	// subtype
	assert.Equal(t, uint8(Original), data[10])

	// command id
	assert.Equal(t, uint8(0), data[11])
	assert.Equal(t, uint8(0x50), data[12])

	// payload
	// TODO
}

func TestReliablePeerInit(t *testing.T) {
	p := CreateReliable(2, commands.NewClientPeerInit())
	p.SeqNr = 65500
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
