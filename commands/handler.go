package commands

type ServerCommandHandler interface {
	OnServerSetPeer(peer *ServerSetPeer)
	OnServerHello(hello *ServerHello)
	OnServerSRPBytesSB(bytesSB *ServerSRPBytesSB)
	OnServerAuthAccept(auth *ServerAuthAccept)
	OnServerAnnounceMedia(announce *ServerAnnounceMedia)
	OnServerCSMRestrictionFlags(flags *ServerCSMRestrictionFlags)
	OnServerBlockData(block *ServerBlockData)
	OnServerTimeOfDay(tod *ServerTimeOfDay)
	OnServerChatMessage(msg *ServerChatMessage)
}
