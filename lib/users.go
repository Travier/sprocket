package lib

import (
	"net"

	xid "github.com/rs/xid"
)

type User struct {
	ID         xid.Id
	Nick       string
	Connection net.Conn
}
