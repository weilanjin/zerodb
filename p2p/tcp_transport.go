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

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc // 🤝握手器
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	Listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		peers:            map[net.Addr]Peer{},
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.Listener, err = net.Listen("tcp", t.ListenAddr)
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

func (t *TCPTransport) handleConnection(conn net.Conn) {
	defer conn.Close()
	peer := NewTCPPeer(conn, true)
	if err := t.HandshakeFunc(peer); err != nil {
		log.Println("Tcp handshake error: ", err)
		return
	}

	// Read loop
	msg := &Message{}
	// buf := make([]byte, 1024)
	for {
		// n, err := conn.Read(buf)
		// if err != nil {
		// 	log.Println("Tcp read error: ", err)
		// 	continue
		// }
		if err := t.Decoder.Decode(conn, msg); err != nil {
			log.Println("Error decoding message:", err)
			continue
		}
		msg.From = conn.RemoteAddr()
		log.Printf("Received message: %+v\n", msg)
	}
}
