package commands

const (
	ClientCommandPeerInit  uint16 = 0
	ClientCommandInit      uint16 = 2
	ClientCommandFirstSRP  uint16 = 80
	ClientCommandSRPBytesA uint16 = 81
	ClientCommandSRPBytesM uint16 = 82
)

const (
	ServerCommandSetPeer               uint16 = 1
	ServerCommandHello                 uint16 = 2
	ServerCommandAccessDenied          uint16 = 10
	ServerCommandTimeOfDay             uint16 = 41
	ServerCommandChatMessage           uint16 = 47
	ServerCommandDetachedInventory     uint16 = 67
	ServerCommandDeleteParticleSpawner uint16 = 83
	ServerCommandUpdatePlayerList      uint16 = 86
	ServerCommandSRPBytesSB            uint16 = 96
)
