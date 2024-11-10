package commands

import (
	"encoding/binary"
)

func GetCommandID(payload []byte) uint16 {
	return binary.BigEndian.Uint16(payload[0:])
}

func GetCommandPayload(payload []byte) []byte {
	return payload[2:]
}

func Parse(payload []byte) (Command, error) {
	commandId := GetCommandID(payload)
	commandPayload := GetCommandPayload(payload)

	var cmd Command

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
	case ServerCommandNodeDefinitions:
		cmd = &ServerNodeDefinitions{}
	}

	if cmd != nil {
		err := cmd.UnmarshalPacket(commandPayload)
		return cmd, err
	}

	return nil, nil
}
