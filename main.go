package main

import (
	"fmt"
	"log"

	"github.com/weilanjin/zerodb/p2p"
)

func OnPeer(peer p2p.Peer) error {
	log.Println("doing some logic with the peer outside of TCPTransport")
	peer.Close()
	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshake,
		Decoder:       &p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for rpc := range tr.Consume() {
			fmt.Println("Got RPC:", rpc)
		}
	}()
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("We DB!")
	select {}
}
