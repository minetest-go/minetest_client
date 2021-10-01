package commands

import (
	"encoding/binary"
	"fmt"
)

type ClientReady struct {
	VersionMajor    uint8
	VersionMinor    uint8
	VersionPatch    uint8
	FullVersion     string
	FormspecVersion uint16
}

func NewClientReady(maj, min, patch uint8, version string, formspecversion uint16) *ClientReady {
	return &ClientReady{
		VersionMajor:    maj,
		VersionMinor:    min,
		VersionPatch:    patch,
		FullVersion:     version,
		FormspecVersion: formspecversion,
	}
}

func (p *ClientReady) GetCommandId() uint16 {
	return ClientCommandReady
}

func (p *ClientReady) MarshalPacket() ([]byte, error) {
	packet := make([]byte, 3+2+1)
	packet[0] = byte(p.VersionMajor)
	packet[1] = byte(p.VersionMinor)
	packet[2] = byte(p.VersionPatch)

	version_length := uint16(len(p.FullVersion))
	binary.BigEndian.PutUint16(packet[4:], version_length)
	packet = append(packet, []byte(p.FullVersion)...)
	packet = append(packet, make([]byte, 2)...)
	binary.BigEndian.PutUint16(packet[6+version_length:], p.FormspecVersion)
	return packet, nil
}

func (p *ClientReady) UnmarshalPacket([]byte) error {
	return nil
}

func (p *ClientReady) String() string {
	return fmt.Sprintf("{ClientReady Version=%d.%d.%d, FullVersion=%s, FormspecVersion=%d}",
		p.VersionMajor, p.VersionMinor, p.VersionPatch, p.FullVersion, p.FormspecVersion)
}
