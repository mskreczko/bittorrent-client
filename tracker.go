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
	AnnounceAction          uint32 = 1
)

func GetPeersList(url string, torrent BitTorrent) {
	socket, err := net.ResolveUDPAddr("udp4", url)
	if err != nil {
		fmt.Println(err)
	}
	conn, err := net.DialUDP("udp4", nil, socket)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	connId := obtainConnectionId(conn)
	fmt.Println(announce(conn, connId, torrent.info))
}

func obtainConnectionId(conn *net.UDPConn) uint64 {
	buf := make([]byte, 16)
	transactionId := getTransactionId()
	binary.BigEndian.PutUint64(buf[0:8], ProtocolId)             // protocol_id
	binary.BigEndian.PutUint32(buf[8:], RequestConnectionAction) // action
	binary.BigEndian.PutUint32(buf[12:], transactionId)          // transaction_id
	n, err := conn.Write(buf)
	if err != nil {
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

func announce(conn *net.UDPConn, connId uint64, info Info) uint32 {
	buf := make([]byte, 98)
	transactionId := getTransactionId()
	binary.BigEndian.PutUint64(buf[0:8], connId)                               // connection_id
	binary.BigEndian.PutUint32(buf[8:12], AnnounceAction)                      // action
	binary.BigEndian.PutUint32(buf[12:16], transactionId)                      // transaction_id
	binary.BigEndian.PutUint32(buf[16:36], binary.BigEndian.Uint32(info.hash)) // info_hash
	binary.BigEndian.PutUint32(buf[36:56], getPeerId())                        // peer_id
	binary.BigEndian.PutUint64(buf[56:64], 0)                                  // downloaded
	binary.BigEndian.PutUint64(buf[64:72], uint64(info.length))                // left
	binary.BigEndian.PutUint64(buf[72:80], 0)                                  // uploaded
	binary.BigEndian.PutUint16(buf[96:], 1337)                                 // port

	n, err := conn.Write(buf)
	if err != nil || n != len(buf) {
		fmt.Println(err)
	}

	n, err = conn.Read(buf)

	if n < 20 {
		fmt.Printf("Response has invalid size: %d\n", n)
	}

	if responseTransactionId := binary.BigEndian.Uint32(buf[4:8]); responseTransactionId != transactionId {
		fmt.Printf("Response has different transactionId: %d != %d\n", responseTransactionId, transactionId)
	}

	if responseAction := binary.BigEndian.Uint32(buf[:4]); responseAction != AnnounceAction {
		fmt.Printf("Response has different action: %d != %d\n", responseAction, AnnounceAction)
	}

	return binary.BigEndian.Uint32(buf[8:12])
}

func getTransactionId() uint32 {
	return uint32(rand.Int31())
}

func getPeerId() uint32 {
	return uint32(rand.Int31())
}
