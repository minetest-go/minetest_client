package commands

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
)

type NodeDefinition struct {
	ID      uint16
	Version uint8
	Name    string
}

func (def *NodeDefinition) Parse(data []byte) error {
	def.Version = data[0]

	offset := 1
	data_len := binary.BigEndian.Uint16(data[offset:])
	offset += 2
	def.Name = string(data[offset : offset+(int(data_len))])
	offset += int(data_len)

	//TODO: more fields
	return nil
}

func (def *NodeDefinition) String() string {
	return fmt.Sprintf("{NodeDefinition ID=%d,Version=%d,Name=%s}", def.ID, def.Version, def.Name)
}

type ServerNodeDefinitions struct {
	Version     uint8
	Count       uint16
	Definitions []*NodeDefinition
}

func (p *ServerNodeDefinitions) GetCommandId() uint16 {
	return ServerCommandNodeDefinitions
}

func (p *ServerNodeDefinitions) MarshalPacket() ([]byte, error) {
	return nil, nil
}

func (p *ServerNodeDefinitions) UnmarshalPacket(payload []byte) error {
	p.Definitions = make([]*NodeDefinition, 0)
	r, err := zlib.NewReader(bytes.NewReader(payload[4:]))
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer([]byte{})
	_, err = io.Copy(buf, r)
	if err != nil {
		return err
	}

	raw_nodes := buf.Bytes()
	p.Version = raw_nodes[0]
	p.Count = binary.BigEndian.Uint16(raw_nodes[1:])

	//fmt.Printf("Nodedefs: version=%d, count=%d\n", p.Version, p.Count)

	nodedefs_size := binary.BigEndian.Uint32(raw_nodes[3:])
	nodedefs_raw := raw_nodes[7 : 7+nodedefs_size]

	offset := 0
	for i := 0; i < int(p.Count); i++ {
		nodeid := binary.BigEndian.Uint16(nodedefs_raw[offset:])
		offset += 2

		nodedef_size := binary.BigEndian.Uint16(nodedefs_raw[offset:])
		offset += 2
		nodedef_raw := nodedefs_raw[offset : offset+int(nodedef_size)]
		nodedef := &NodeDefinition{ID: nodeid}
		nodedef.Parse(nodedef_raw)

		p.Definitions = append(p.Definitions, nodedef)

		offset += int(nodedef_size)
	}

	return nil
}

func (p *ServerNodeDefinitions) String() string {
	return fmt.Sprintf("{ServerNodeDefinitions Version=%d, Count=%d, Definitions=%s}", p.Version, p.Count, p.Definitions)
}
