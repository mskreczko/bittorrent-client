package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
)

const (
	ProtocolId              uint64 = 0x41727101980
	RequestConnectionAction uint32 = 0
)

func GetPeersList(url string) {
	socket, err := net.ResolveUDPAddr("udp4", url)
	if err != nil {
		fmt.Println(err)
	}
	conn, err := net.DialUDP("udp4", nil, socket)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	fmt.Println(ObtainConnectionId(conn))
}

func ObtainConnectionId(conn *net.UDPConn) uint64 {
	buf := make([]byte, 16)
	transactionId := uint32(rand.Int31())
	binary.BigEndian.PutUint64(buf[0:8], ProtocolId)
	binary.BigEndian.PutUint32(buf[8:], RequestConnectionAction)
	binary.BigEndian.PutUint32(buf[12:], transactionId)
	n, err := conn.Write(buf)
	if err != nil || n != len(buf) {
		fmt.Println(err)
	}

	n, err = conn.Read(buf)

	if n != len(buf) {
		fmt.Printf("Response has different size: %d != %d\n", n, len(buf))
	}
	if responseTransactionId := binary.BigEndian.Uint32(buf[4:8]); responseTransactionId != transactionId {
		fmt.Printf("Response has different transactionId: %d != %d\n", responseTransactionId, transactionId)
	}

	if responseAction := binary.BigEndian.Uint32(buf[:4]); responseAction != RequestConnectionAction {
		fmt.Printf("Response has different action: %d != %d\n", responseAction, RequestConnectionAction)
	}

	return binary.BigEndian.Uint64(buf[8:])
}
