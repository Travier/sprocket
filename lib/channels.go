package lib

import (
	"fmt"
	"strconv"
)

type Channel struct {
	Name        string
	Messages    []string
	Connections []TCPConnection
}

func (channel Channel) HasConnection(conn TCPConnection) bool {
	for i := range channel.Connections {
		chanConn := channel.Connections[i]

		if chanConn.ID == conn.ID {
			return true
		}
	}

	return false
}

func SendChannelMessage(channel Channel, message string) {
	fmt.Println("Sending message to " + strconv.Itoa(len(channel.Connections)))
	for i := range channel.Connections {
		conn := channel.Connections[i]

		conn.SendMessage(message)
	}
}

func JoinChannel(list []Channel, channelName string, conn TCPConnection) bool {
	for i := range list {
		channel := &list[i]

		if channelName == channel.Name {
			if !channel.HasConnection(conn) {
				//channel doesn't have this connection yet
				channel.Connections = append(channel.Connections, conn)

				conn.SendMessage("WELCOME TO THE CHANNEL")
				return true
			}
		}
	}

	return false
}

func CreateChannel(name string) Channel {
	channel := Channel{Name: name}

	return channel
}
