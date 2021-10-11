package packet

import (
	"fmt"
	"minetest_client/commands"
)

func CreateCommand(commandId uint16, payload []byte) (Command, error) {
	var cmd Command = nil
	switch commandId {
	case commands.ServerCommandSetPeer:
		cmd = &commands.ServerSetPeer{}
	case commands.ServerCommandHello:
		cmd = &commands.ServerHello{}
	case commands.ServerCommandAuthAccept:
		cmd = &commands.ServerAuthAccept{}
	case commands.ServerCommandAccessDenied:
		cmd = &commands.ServerAccessDenied{}
	case commands.ServerCommandTimeOfDay:
		cmd = &commands.ServerTimeOfDay{}
	case commands.ServerCommandChatMessage:
		cmd = &commands.ServerChatMessage{}
	case commands.ServerCommandDetachedInventory:
		cmd = &commands.ServerDetachedInventory{}
	case commands.ServerCommandDeleteParticleSpawner:
		cmd = &commands.ServerDeleteParticleSpawner{}
	case commands.ServerCommandUpdatePlayerList:
		cmd = &commands.ServerUpdatePlayerList{}
	case commands.ServerCommandSRPBytesSB:
		cmd = &commands.ServerSRPBytesSB{}
	case commands.ServerCommandItemDefinitions:
		cmd = &commands.ServerItemDefinitions{}
	case commands.ServerCommandNodeDefinitions:
		cmd = &commands.ServerNodeDefinitions{}
	case commands.ServerCommandAnnounceMedia:
		cmd = &commands.ServerAnnounceMedia{}
	case commands.ServerCommandMovement:
		cmd = &commands.ServerMovement{}
	case commands.ServerCommandCSMRestrictionFlags:
		cmd = &commands.ServerCSMRestrictionFlags{}
	case commands.ServerCommandBlockData:
		cmd = &commands.ServerBlockData{}
	case commands.ServerCommandHudAdd:
		cmd = &commands.ServerHudAdd{}
	case commands.ServerCommandHudSetParam:
		cmd = &commands.ServerHudSetParam{}
	case commands.ServerCommandHudSetFlags:
		cmd = &commands.ServerHudSetFlags{}
	case commands.ServerCommandHudChange:
		cmd = &commands.ServerHudChange{}
	case commands.ServerCommandActiveObjectMessage:
		//TODO
	case commands.ServerCommandAddParticleSpawner:
		//TODO
	default:
		fmt.Printf("Unknown command received: %d\n", commandId)
	}

	if cmd != nil {
		//fmt.Printf("Unmarshal: %d\n", commandId)
		err := cmd.UnmarshalPacket(payload)
		return cmd, err
	}

	return nil, nil
}
