package commands

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerHello(t *testing.T) {
	f, err := os.Open("testdata/server_hello.bin")
	assert.NoError(t, err)
	assert.NotNil(t, f)
	defer f.Close()

	fmt.Print(f)

	payload := make([]byte, 64)
	count, err := f.Read(payload)
	assert.NoError(t, err)
	assert.True(t, count > 0)

	pkg := &ServerHello{}
	err = pkg.UnmarshalPacket(payload[0x37:])
	assert.NoError(t, err)

	assert.Equal(t, uint8(28), pkg.SerializationVersion)
}
