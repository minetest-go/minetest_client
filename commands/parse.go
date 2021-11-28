package commands

import (
	"encoding/binary"
	"fmt"
	"minetest_client/packet"
)

func Parse(payload []byte) (packet.Command, error) {
	commandId := binary.BigEndian.Uint16(payload[0:])
	commandPayload := payload[2:]

	var cmd packet.Command

	switch commandId {
	case ServerCommandSetPeer:
		cmd = &ServerSetPeer{}
	case ServerCommandHello:
		cmd = &ServerHello{}
	case ServerCommandSRPBytesSB:
		cmd = &ServerSRPBytesSB{}
	case ServerCommandAuthAccept:
		cmd = &ServerAuthAccept{}
	case ServerCommandAnnounceMedia:
		cmd = &ServerAnnounceMedia{}
	case ServerCommandCSMRestrictionFlags:
		cmd = &ServerCSMRestrictionFlags{}
	case ServerCommandBlockData:
		cmd = &ServerBlockData{}
	case ServerCommandTimeOfDay:
		cmd = &ServerTimeOfDay{}
	case ServerCommandChatMessage:
		cmd = &ServerChatMessage{}
	case ServerCommandAddParticleSpawner:
		cmd = &ServerAddParticleSpawner{}
	case ServerCommandDetachedInventory:
		cmd = &ServerDetachedInventory{}
	case ServerCommandHudChange:
		cmd = &ServerHudChange{}
	case ServerCommandActiveObjectMessage:
		cmd = &ServerActiveObjectMessage{}
	case ServerCommandDeleteParticleSpawner:
		cmd = &ServerDeleteParticleSpawner{}
	case ServerCommandMovePlayer:
		cmd = &ServerMovePlayer{}
	case ServerCommandMedia:
		cmd = &ServerMedia{}
	case ServerCommandAccessDenied:
		cmd = &ServerAccessDenied{}
	default:
		fmt.Printf("Unknown command received: %d\n", commandId)
	}

	if cmd != nil {
		err := cmd.UnmarshalPacket(commandPayload)
		return cmd, err
	}

	return nil, nil
}
