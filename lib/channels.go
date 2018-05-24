package lib

type Channel struct {
	Name        string
	Messages    []string
	Connections []TCPConnection
}

func (channel Channel) GlobalMessage(message string) {
	for i := range channel.Connections {
		conn := channel.Connections[i]

		conn.SendMessage(message)
	}
}

func (channel *Channel) Join(conn TCPConnection) {
	if !channel.HasConnection(conn) {
		channel.Connections = append(channel.Connections, conn)
	}
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

func CreateChannel(name string) Channel {
	channel := Channel{Name: name}

	return channel
}
