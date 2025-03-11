package p2p

func NOPHandshake(Peer) error {
	return nil
}

type HandshakeFunc func(Peer) error
