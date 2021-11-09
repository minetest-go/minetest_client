package commands

import (
	"encoding/binary"
	"fmt"
	"sort"
)

type ClientRequestMedia struct {
	Names []string
}

func NewClientRequestMedia(names []string) *ClientRequestMedia {
	return &ClientRequestMedia{
		Names: names,
	}
}

func (p *ClientRequestMedia) GetCommandId() uint16 {
	return ClientCommandRequestMedia
}

func (p *ClientRequestMedia) MarshalPacket() ([]byte, error) {
	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, uint16(len(p.Names)))

	sort.Strings(p.Names)
	for _, name := range p.Names {
		name_len := make([]byte, 2)
		binary.BigEndian.PutUint16(name_len, uint16(len(name)))

		data = append(data, name_len...)
		data = append(data, []byte(name)...)
	}

	return data, nil
}

func (p *ClientRequestMedia) UnmarshalPacket([]byte) error {
	return nil
}

func (p *ClientRequestMedia) String() string {
	return fmt.Sprintf("{ClientRequestMedia #Names=%d}", len(p.Names))
}
