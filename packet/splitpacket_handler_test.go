package packet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitPacketHandler(t *testing.T) {
	sph := NewSplitPacketHandler()

	res := sph.AddPacket(&SplitPayload{
		SeqNr:       1,
		ChunkCount:  2,
		ChunkNumber: 0,
		Data:        []byte{0x00, 0x01},
	})

	assert.Nil(t, res)

	res = sph.AddPacket(&SplitPayload{
		SeqNr:       1,
		ChunkCount:  2,
		ChunkNumber: 1,
		Data:        []byte{0x02, 0x03},
	})

	assert.NotNil(t, res)
	assert.Equal(t, 4, len(res))
	assert.Equal(t, uint8(0x00), res[0])
	assert.Equal(t, uint8(0x01), res[1])
	assert.Equal(t, uint8(0x02), res[2])
	assert.Equal(t, uint8(0x03), res[3])
}
