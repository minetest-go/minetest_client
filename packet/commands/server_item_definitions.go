package commands

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
)

type ItemType uint8

const (
	ITEM_NONE  ItemType = 0
	ITEM_NODE  ItemType = 1
	ITEM_CRAFT ItemType = 2
	ITEM_TOOL  ItemType = 3
)

func (it ItemType) String() string {
	switch it {
	case ITEM_NONE:
		return "None"
	case ITEM_NODE:
		return "Node"
	case ITEM_CRAFT:
		return "Craft"
	case ITEM_TOOL:
		return "Tool"
	}

	return "Unknown"
}

type ItemDefinition struct {
	Version     uint8
	Type        ItemType
	Name        string
	Description string
}

func (def *ItemDefinition) Parse(data []byte) error {
	def.Version = data[0]
	def.Type = ItemType(data[1])

	offset := 2
	name_len := binary.BigEndian.Uint16(data[offset:])
	offset += 2
	def.Name = string(data[offset : offset+int(name_len)])
	offset += int(name_len)
	desc_len := binary.BigEndian.Uint16(data[offset:])
	offset += 2
	def.Description = string(data[offset : offset+int(desc_len)])

	//TODO: more fields
	return nil
}

func (def *ItemDefinition) String() string {
	return fmt.Sprintf("{ItemDefinition Version=%d, Type=%s, Name='%s', Description='%s'}",
		def.Version, def.Type, def.Name, def.Description)
}

type ServerItemDefinitions struct {
	Version     uint8
	Count       uint16
	Definitions []*ItemDefinition
}

func (p *ServerItemDefinitions) GetCommandId() uint16 {
	return ServerCommandItemDefinitions
}

func (p *ServerItemDefinitions) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerItemDefinitions) UnmarshalPacket(payload []byte) error {
	p.Definitions = make([]*ItemDefinition, 0)
	r, err := zlib.NewReader(bytes.NewReader(payload))
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer([]byte{})
	_, err = io.Copy(buf, r)
	if err != nil {
		return err
	}

	raw_items := buf.Bytes()
	p.Version = raw_items[0]
	p.Count = binary.BigEndian.Uint16(raw_items[1:])
	offset := 3
	for i := 0; i < int(p.Count); i++ {
		itemdef_size := binary.BigEndian.Uint16(raw_items[offset:])
		offset += 2
		itemdef_raw := raw_items[offset : offset+int(itemdef_size)]
		itemdef := &ItemDefinition{}
		itemdef.Parse(itemdef_raw)

		p.Definitions = append(p.Definitions, itemdef)

		offset += int(itemdef_size)
	}

	// TODO: aliases
	//fmt.Println(buf.String())

	return nil
}

func (p *ServerItemDefinitions) String() string {
	return fmt.Sprintf("{ServerItemDefinitions Version=%d, Count=%d, Definitions=%s}", p.Version, p.Count, p.Definitions)
}
