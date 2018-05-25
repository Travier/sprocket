package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strconv"

	. "github.com/Travier/sprocket/lib"
	"github.com/rs/xid"
)

var addr = flag.String("addr", "", "The address to listen to; default is \"\" (all interfaces).")
var port = flag.Int("port", 9999, "The port to listen on; default is 9999.")
var channelList = make([]Channel, 1)
var userList = make([]User, 0)
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

		go handleConnection(connection)
	}
}

func handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)

	user := *User{id: xid.New(), Nick: "", Connection: conn}
	userList = append(userList, user)

	mainChan.Join(conn)

	scanner := bufio.NewScanner(conn.Instance)

	for {
		ok := scanner.Scan()

		handleMessage(user, scanner.Text())

		if !ok {
			break
		}
	}

	fmt.Println("Client at " + remoteAddr + " disconnected.")
}

func handleMessage(user *User, message string) {

	if isCommand("/nick", message) {

	}
}

func isCommand(command string, input) bool {
	if input.Contains(command) {
		return true
	}

	return false
}
