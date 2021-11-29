package commands

import "encoding/binary"

func CreatePayload(cmd Command) ([]byte, error) {
	inner_payload, err := cmd.MarshalPacket()
	if err != nil {
		return nil, err
	}

	payload := make([]byte, len(inner_payload)+2)
	copy(payload[2:], inner_payload)

	binary.BigEndian.PutUint16(payload[0:], cmd.GetCommandId())

	return payload, nil
}
