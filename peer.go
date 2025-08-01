package main

import (
	"net"
)

type Peer struct {
	conn  net.Conn
	msgCh chan Message // Channel to send message to server
}

func NewPeer(conn net.Conn, msgCh chan Message) *Peer {
	return &Peer{
		conn:  conn,
		msgCh: msgCh,
	}
}

func (p *Peer) Send(msg []byte) (int, error) {
	return p.conn.Write(msg)
}

func (p *Peer) readLoop() error {
	buf := make([]byte, 1024)
	for {
		n, err := p.conn.Read(buf)
		if err != nil {
			return err
		}

		// Creating new buffer for each message (preventing buffer race)
		msgBuf := make([]byte, n)
		copy(msgBuf, buf[:n])
		p.msgCh <- Message{
			data: msgBuf,
			peer: p,
		}
	}
}
