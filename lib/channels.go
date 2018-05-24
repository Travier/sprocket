package channels

import "net"

type Channel struct {
	name        string
	messages    []string
	connections []net.Conn
}

func createChannel(name string) Channel {
	channel := Channel{name: name}

	return channel
}
