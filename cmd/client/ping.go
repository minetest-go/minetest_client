package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"time"
)

var INIT_PKG = []byte{0x4f, 0x45, 0x74, 0x03, 0x00, 0x00, 0x00, 0x01}

func createDiscoPkg(peer_id uint16) []byte {
	pkg := []byte{0x4f, 0x45, 0x74, 0x03, 0xff, 0xff, 0x00, 0x00, 0x03}
	binary.BigEndian.PutUint16(pkg[4:], peer_id)

	return pkg
}

type PingResult struct {
	Delay  time.Duration
	PeerID uint16
}

func Ping(host string, port int) (*PingResult, error) {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	start := time.Now()

	conn.Write(INIT_PKG)

	buf := make([]byte, 128)
	count, err := conn.Read(buf)
	if err != nil {
		return nil, err

	} else if count == 0 {
		return nil, errors.New("no data received")

	} else if count != 14 {
		return nil, fmt.Errorf("invalid packet received: len=%d", count)

	}

	res := &PingResult{
		Delay:  time.Now().Sub(start),
		PeerID: binary.BigEndian.Uint16(buf[12:]),
	}

	disco := createDiscoPkg(res.PeerID)
	_, err = conn.Write(disco)
	if err != nil {
		return nil, err
	}

	return res, nil
}
