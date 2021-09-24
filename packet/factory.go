package packet

import "minetest_client/packet/commands"

func CreateCommand(commandId uint16, payload []byte) (Command, error) {
	var cmd Command = nil
	switch commandId {
	case 1:
		cmd = &commands.ServerSetPeer{}
	case 2:
		cmd = &commands.ServerHello{}
	case 41:
		cmd = &commands.ServerTimeOfDay{}
	case 47:
		cmd = &commands.ServerChatMessage{}
	case 67:
		cmd = &commands.ServerDetachedInventory{}
	case 83:
		cmd = &commands.ServerDeleteParticleSpawner{}
	case 86:
		cmd = &commands.ServerUpdatePlayerList{}
	}

	if cmd != nil {
		err := cmd.UnmarshalPacket(payload)
		return cmd, err
	}

	return nil, nil
}
