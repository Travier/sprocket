package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strconv"

	. "github.com/Travier/sprocket/lib"
	xid "github.com/rs/xid"
)

var addr = flag.String("addr", "", "The address to listen to; default is \"\" (all interfaces).")
var port = flag.Int("port", 9999, "The port to listen on; default is 9999.")
var channelList = make([]Channel, 1)
var mainChan = CreateChannel("main")

func main() {
	fmt.Println("Starting server...")

	src := *addr + ":" + strconv.Itoa(*port)
	listener, _ := net.Listen("tcp", src)
	fmt.Printf("Listening on %s.\n", src)

	//create 'main' channel

	//add channel to list
	channelList = append(channelList, mainChan)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Some connection error: %s\n", err)
		}

		connection := TCPConnection{ID: xid.New(), Instance: conn}

		go handleConnection(connection)
	}
}

func handleConnection(conn TCPConnection) {
	remoteAddr := conn.Instance.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)

	mainChan.Join(conn)

	scanner := bufio.NewScanner(conn.Instance)

	for {
		ok := scanner.Scan()

		handleMessage(conn, scanner.Text())

		if !ok {
			break
		}
	}

	fmt.Println("Client at " + remoteAddr + " disconnected.")
}

func handleMessage(conn TCPConnection, message string) {
	if message[0] == '/' {
		switch {
		case message == "/motd":
			//SendMessage(conn, "The server is running great today! I wonder if longer texts makes all the difference here prolly but idk")
		case message == "/time":
		//	resp := "It is " + time.Now().String() + "\n"
		//SendMessage(conn, resp)
		default:
			//SendMessage(conn, "Unrecognized command.")
		}
	} else {
		mainChan.GlobalMessage(message)
	}
}
