package commands

type ServerCommandHandler interface {
	OnServerSetPeer(peer *ServerSetPeer)
	OnServerHello(hello *ServerHello)
	OnServerSRPBytesSB(bytesSB *ServerSRPBytesSB)
	OnServerAuthAccept(auth *ServerAuthAccept)
	OnServerAnnounceMedia(announce *ServerAnnounceMedia)
	OnServerMedia(media *ServerMedia)
	OnServerCSMRestrictionFlags(flags *ServerCSMRestrictionFlags)
	OnServerBlockData(block *ServerBlockData)
	OnServerTimeOfDay(tod *ServerTimeOfDay)
	OnServerChatMessage(msg *ServerChatMessage)
	OnServerMovePlayer(msg *ServerMovePlayer)
	OnAddParticleSpawner(aps *ServerAddParticleSpawner)
	OnDeleteParticleSpawner(msg *ServerDeleteParticleSpawner)
	OnHudChange(hud *ServerHudChange)
	OnDetachedInventory(inv *ServerDetachedInventory)
	OnActiveObjectMessage(aom *ServerActiveObjectMessage)
}
