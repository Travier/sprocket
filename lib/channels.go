package lib

import "net"

type Channel struct {
	name        string
	messages    []string
	connections []net.Conn
}

func JoinChannel(list []Channel) {

}

func CreateChannel(name string) Channel {
	channel := Channel{name: name}

	return channel
}
