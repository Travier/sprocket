package lib

import (
	"net"

	"github.com/rs/xid"
)

type TCPConnection struct {
	ID       xid.ID
	Instance net.Conn
}

func SendMessage(conn TCPConnection, message string) {
	conn.Instance.Write([]byte(message + "\n"))
}
