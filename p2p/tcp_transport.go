package p2p

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// TCPeer represents the remote node over a TCP established connection.
type TCPPeer struct {
	// conn is the underlying connection of the peer
	conn net.Conn

	// if we dial and retrieve a comm => outbound = true
	// if we accept and retrieve a comm => outbound = false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	ListenAddress string
	Listener      net.Listener
	shakeHands    HandshakeFunc // 🤝握手器
	decoder       Decoder
	mu            sync.RWMutex
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		ListenAddress: listenAddr,
		shakeHands:    NOPHandshake,
		peers:         map[net.Addr]Peer{},
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.Listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		fmt.Printf("new incoming connection %+v\n", conn)
		go t.handleConnection(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	defer conn.Close()
	peer := NewTCPPeer(conn, true)
	if err := t.shakeHands(peer); err != nil {
	}

	// Read loop
	msg := &Temp{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			log.Println("Error decoding message:", err)
			continue
		}
	}
}
