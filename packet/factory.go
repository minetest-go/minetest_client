package packet

import "minetest_client/packet/commands"

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
	}

	if cmd != nil {
		err := cmd.UnmarshalPacket(payload)
		return cmd, err
	}

	return nil, nil
}
