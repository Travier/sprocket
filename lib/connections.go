package lib

import (
	"net"

	"github.com/rs/xid"
)

type TCPConnection struct {
	ID       xid.ID
	Instance net.Conn
}

func (conn TCPConnection) SendMessage(message string) {
	conn.Instance.Write([]byte(message + "\n"))
}
