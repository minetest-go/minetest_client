package commands

// from Client to Server
const (
	ClientCommandPeerInit  uint16 = 0
	ClientCommandInit      uint16 = 2
	ClientCommandInit2     uint16 = 17
	ClientCommandPlayerPos uint16 = 35
	ClientCommandGotBlocks uint16 = 36
	ClientCommandReady     uint16 = 67
	ClientCommandFirstSRP  uint16 = 80
	ClientCommandSRPBytesA uint16 = 81
	ClientCommandSRPBytesM uint16 = 82
)

// from Server to Client
const (
	ServerCommandSetPeer               uint16 = 1
	ServerCommandHello                 uint16 = 2
	ServerCommandAuthAccept            uint16 = 3
	ServerCommandAccessDenied          uint16 = 10
	ServerCommandBlockData             uint16 = 32
	ServerCommandTimeOfDay             uint16 = 41
	ServerCommandCSMRestrictionFlags   uint16 = 42
	ServerCommandChatMessage           uint16 = 47
	ServerCommandNodeDefinitions       uint16 = 58
	ServerCommandAnnounceMedia         uint16 = 60
	ServerCommandItemDefinitions       uint16 = 61
	ServerCommandDetachedInventory     uint16 = 67
	ServerCommandMovement              uint16 = 69
	ServerCommandHudAdd                uint16 = 73
	ServerCommandHudChange             uint16 = 75
	ServerCommandHudSetFlags           uint16 = 76
	ServerCommandHudSetParam           uint16 = 77
	ServerCommandDeleteParticleSpawner uint16 = 83
	ServerCommandUpdatePlayerList      uint16 = 86
	ServerCommandSRPBytesSB            uint16 = 96
)
