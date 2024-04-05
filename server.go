package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Server running...")
	server, err := net.Listen("tcp", "localhost:6969")
	if err != nil {
		panic(err)
	}
	defer server.Close()

	for {
		connection, err := server.Accept()
		if err != nil {
			panic(err)
		}
		go clientprocess(connection)
	}
}

func clientprocess(conn net.Conn) {

	//read message
	var size uint32
	err := binary.Read(conn, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}
  err = conn.SetReadDeadline(time.Now().Add(3 * time.Second))
  if err != nil{
    panic(err)
  }

	buffer := make([]byte, size)
	_, err = conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	msg := string(buffer)
	fmt.Println("Recived: " + msg)

	// write replay
	var replay string
	if strings.HasSuffix(msg, ".zip") {
		replay = "file recived"
	} else if strings.HasSuffix(msg, ".") {
		replay = "only zip are allowed"
	} else {
		replay = "massage recived"
	}

  err = conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
  if err != nil{
    panic(err)
  }
	err = binary.Write(conn, binary.LittleEndian, uint32(len(replay)))
	if err != nil {
		panic(err)
	}
	_, err = conn.Write([]byte(replay))
	if err != nil {
		panic(err)
	}
}
