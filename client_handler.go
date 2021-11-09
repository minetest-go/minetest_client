package main

import (
	"fmt"
	"minetest_client/commands"
	"minetest_client/packet"
	"minetest_client/srp"
	"time"
)

type ClientHandler struct {
	Client      *Client
	Username    string
	Password    string
	StalkMode   bool
	SRPPubA     []byte
	SRPPrivA    []byte
	MediaHashes map[string][]byte
}

func (ch *ClientHandler) OnServerSetPeer(peer *commands.ServerSetPeer) {
	fmt.Printf("Received set_peerid: %d\n", peer.PeerID)
	go func() {
		time.Sleep(1 * time.Second)
		err := ch.Client.SendOriginalCommand(commands.NewClientInit(ch.Username))
		if err != nil {
			panic(err)
		}
	}()
}

func (ch *ClientHandler) OnServerHello(hello *commands.ServerHello) {
	if ch.StalkMode {
		return
	}
	packet.ResetSeqNr(65500)

	if hello.AuthMechanismSRP {
		// existing client
		var err error
		ch.SRPPubA, ch.SRPPrivA, err = srp.InitiateHandshake()
		if err != nil {
			panic(err)
		}

		fmt.Println("Sending SRP bytes A")
		err = ch.Client.SendCommand(commands.NewClientSRPBytesA(ch.SRPPubA))
		if err != nil {
			panic(err)
		}
	}

	if hello.AuthMechanismFirstSRP {
		// new client
		salt, verifier, err := srp.NewClient([]byte(ch.Username), []byte(ch.Password))
		if err != nil {
			panic(err)
		}

		fmt.Println("Sending first SRP")
		err = ch.Client.SendCommand(commands.NewClientFirstSRP(salt, verifier))
		if err != nil {
			panic(err)
		}
	}
}

func (ch *ClientHandler) OnServerSRPBytesSB(bytesSB *commands.ServerSRPBytesSB) {
	identifier := []byte(ch.Username)
	passphrase := []byte(ch.Password)

	clientK, err := srp.CompleteHandshake(ch.SRPPubA, ch.SRPPrivA, identifier, passphrase, bytesSB.BytesS, bytesSB.BytesB)
	if err != nil {
		panic(err)
	}

	proof := srp.ClientProof(identifier, bytesSB.BytesS, ch.SRPPubA, bytesSB.BytesB, clientK)

	fmt.Println("Sending SRP bytes M")
	err = ch.Client.SendCommand(commands.NewClientSRPBytesM(proof))
	if err != nil {
		panic(err)
	}
}

func (ch *ClientHandler) OnServerAuthAccept(auth *commands.ServerAuthAccept) {
	fmt.Println("Sending INIT2")
	err := ch.Client.SendCommand(commands.NewClientInit2())
	if err != nil {
		panic(err)
	}
}

func (ch *ClientHandler) OnServerAnnounceMedia(announce *commands.ServerAnnounceMedia) {
	fmt.Printf("Server announces media: %d assets\n", announce.FileCount)
	ch.MediaHashes = announce.Hashes
}

func (ch *ClientHandler) OnServerMedia(media *commands.ServerMedia) {
	fmt.Printf("Server media: %s\n", media)
}

func (ch *ClientHandler) OnServerCSMRestrictionFlags(flags *commands.ServerCSMRestrictionFlags) {
	fmt.Println("Server sends csm restriction flags")

	fmt.Println("Sending REQUEST_MEDIA")
	files := make([]string, 0)
	for name := range ch.MediaHashes {
		//fmt.Printf("Name: '%s'\n", name)
		files = append(files, name)
	}

	time.Sleep(20 * time.Millisecond)

	reqmedia_cmd := commands.NewClientRequestMedia(files)
	err := ch.Client.SendCommand(reqmedia_cmd)
	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Sending CLIENT_READY")
	err = ch.Client.SendCommand(commands.NewClientReady(5, 5, 5, "mt-bot", 4))
	if err != nil {
		panic(err)
	}

	fmt.Println("Sending PLAYERPOS")
	ppos := commands.NewClientPlayerPos()
	err = ch.Client.SendCommand(ppos)
	if err != nil {
		panic(err)
	}
}

func (ch *ClientHandler) OnServerBlockData(block *commands.ServerBlockData) {
	//fmt.Printf("Block: '%s'\n", block_pkg)

	gotblocks := commands.NewClientGotBlocks()
	gotblocks.AddBlock(block.Pos)

	err := ch.Client.SendCommand(gotblocks)
	if err != nil {
		panic(err)
	}

}

func (ch *ClientHandler) OnServerTimeOfDay(tod *commands.ServerTimeOfDay) {
	fmt.Printf("Time of day: %d\n", tod.TimeOfDay)
}

func (ch *ClientHandler) OnServerChatMessage(msg *commands.ServerChatMessage) {
	fmt.Printf("Chat: '%s'\n", msg.Message)
}

func (ch *ClientHandler) OnAddParticleSpawner(aps *commands.ServerAddParticleSpawner) {

}

func (ch *ClientHandler) OnHudChange(hud *commands.ServerHudChange) {

}

func (ch *ClientHandler) OnDetachedInventory(inv *commands.ServerDetachedInventory) {

}

func (ch *ClientHandler) OnActiveObjectMessage(aom *commands.ServerActiveObjectMessage) {

}

func (ch *ClientHandler) OnDeleteParticleSpawner(msg *commands.ServerDeleteParticleSpawner) {

}

func (ch *ClientHandler) OnServerMovePlayer(msg *commands.ServerMovePlayer) {
	fmt.Printf("Move player: '%s'\n", msg)
}
