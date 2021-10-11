package commands

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strings"
)

type ServerAnnounceMedia struct {
	FileCount     uint16
	Hashes        map[string][]byte
	RemoteServers []string
}

func (p *ServerAnnounceMedia) GetCommandId() uint16 {
	return ServerCommandAnnounceMedia
}

func (p *ServerAnnounceMedia) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerAnnounceMedia) UnmarshalPacket(payload []byte) error {
	p.Hashes = make(map[string][]byte)
	p.RemoteServers = make([]string, 0)

	p.FileCount = binary.BigEndian.Uint16(payload[0:])

	offset := 2
	for i := 0; i < int(p.FileCount); i++ {
		name_len := binary.BigEndian.Uint16(payload[offset:])
		if name_len == 0 {
			return fmt.Errorf("name-len is 0 at %d", offset)
		}
		offset += 2

		name := string(payload[offset : offset+(int(name_len))])
		name = strings.TrimSpace(name)
		if len(name) == 0 {
			return fmt.Errorf("len(name) is 0 at %d", offset)
		}
		offset += int(name_len)

		hash_len := binary.BigEndian.Uint16(payload[offset:])
		offset += 2

		hash := payload[offset : offset+int(hash_len)]
		offset += int(hash_len)

		hash_binary, err := base64.RawStdEncoding.DecodeString(string(hash))
		if err != nil {
			return err
		}

		p.Hashes[name] = hash_binary
	}

	remoteservers_len := binary.BigEndian.Uint16(payload[offset:])
	offset += 2

	remoteservers := payload[offset : offset+int(remoteservers_len)]
	p.RemoteServers = strings.Split(string(remoteservers), ",")

	return nil
}

func (p *ServerAnnounceMedia) String() string {
	return fmt.Sprintf("{ServerAnnounceMedia FileCount=%d, RemoteServers=%s}", p.FileCount, p.RemoteServers)
}
