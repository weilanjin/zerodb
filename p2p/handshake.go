package p2p

// ErrInvalidHandshake is returned if the handshake between
// the local and remote node could not be established.
// var ErrInvalidHandshake = errors.New("invalid handshake")

func NOPHandshake(Peer) error {
	return nil
}

type HandshakeFunc func(Peer) error
