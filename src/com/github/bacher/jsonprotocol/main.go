package main

import (
	"fmt"
	"net"
	"os"
	"encoding/binary"
)

func main() {
	PORT := ":9999"

	server, err := net.Listen("tcp", PORT)

	fmt.Println("Listening started at " + PORT)

	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(2)
	}

	for {
		conn, err := server.Accept()

		if err != nil {
			fmt.Println("Accept failed:", err)
			break
		}

		go handleRequest(conn)
	}

	server.Close()
}

func handleRequest(conn net.Conn) {
	isWaitNewMessage := true
	msgLen := 0
	msgRec := 0
	var msgBuffer []byte = nil

	for {
		if isWaitNewMessage {
			var memRaw uint32 = 0
			err2 := binary.Read(conn, binary.BigEndian, &memRaw)

			if err2 != nil {
				fmt.Println("WHAT IS IT?")
				break
			}
			msgLen = int(memRaw)
			msgRec = 0
		}

		msgBuffer = make([]byte, msgLen)

		reqLen, err := conn.Read(msgBuffer[msgRec:])

		if err != nil {
			fmt.Println("Error reading:", err)
			break
		}

		msgRec += reqLen

		if msgLen == msgRec {
			onReceiveJson(msgBuffer)

			msgBuffer = nil
			isWaitNewMessage = true
			msgLen = 0
			msgRec = 0
		}
	}

	conn.Close()
}

func onReceiveJson(json []byte) {
	fmt.Println("OK -->", string(json))
}
