package commands

type ServerCommandHandler interface {
	OnServerSetPeer(peer *ServerSetPeer)
	OnServerHello(hello *ServerHello)
	OnServerSRPBytesSB(srp *ServerSRPBytesSB)
	OnServerAuthAccept(auth *ServerAuthAccept)
	OnServerCSMRestrictionFlags(flags *ServerCSMRestrictionFlags)
	OnServerBlockData(block *ServerBlockData)
	OnServerTimeOfDay(tod *ServerTimeOfDay)
	OnServerChatMessage(msg *ServerChatMessage)
}
