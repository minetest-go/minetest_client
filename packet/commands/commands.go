package commands

type ClientCommand uint16

const (
	ClientCommandPeerInit  = 0
	ClientCommandInit      = 2
	ClientCommandFirstSRP  = 80
	ClientCommandSRPBytesA = 81
)

type ServerCommand uint16

const (
	ServerCommandSetPeer               = 1
	ServerCommandHello                 = 2
	ServerCommandAccessDenied          = 10
	ServerCommandTimeOfDay             = 41
	ServerCommandChatMessage           = 47
	ServerCommandDetachedInventory     = 67
	ServerCommandDeleteParticleSpawner = 83
	ServerCommandUpdatePlayerList      = 86
	ServerCommandSRPBytesSB            = 96
)
