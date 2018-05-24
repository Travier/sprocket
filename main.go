package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var addr = flag.String("addr", "", "The address to listen to; default is \"\" (all interfaces).")
var port = flag.Int("port", 9999, "The port to listen on; default is 9999.")

func main() {
	fmt.Println("Starting server...")

	src := *addr + ":" + strconv.Itoa(*port)
	listener, _ := net.Listen("tcp", src)
	fmt.Printf("Listening on %s.\n", src)

	channelList := make([]Channel)
	mainChan := createChannel("main")

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Some connection error: %s\n", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)

	scanner := bufio.NewScanner(conn)

	for {
		ok := scanner.Scan()

		handleMessage(scanner.Text(), conn)

		if !ok {
			break
		}
	}

	fmt.Println("Client at " + remoteAddr + " disconnected.")
}

func handleMessage(message string, conn net.Conn) {
	fmt.Println("> " + message)

	if strings.Contains("/channel", message) {

	}

	if message[0] == '/' {
		switch {
		case strings.Contains(message, "/channel"):
			parts := strings.Split(message, " ")
			if len(parts) != 2 {
				break
			}

		case message == "/motd":
			sendMessage(conn, "The server is running great today! I wonder if longer texts makes all the difference here prolly but idk")
		case message == "/time":
			resp := "It is " + time.Now().String() + "\n"
			sendMessage(conn, resp)
		default:
			sendMessage(conn, "Unrecognized command.")
		}
	}
}

func sendMessage(conn net.Conn, message string) {
	conn.Write([]byte(message + "\n"))
}
