package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

func menu() {
	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Send Massage")
		fmt.Println("2. Exit")
		fmt.Print(">> ")
		scan.Scan()
		ch := scan.Text()
		if ch == "1" {
			sendMsgMenu()
		} else if ch == "2" {
			fmt.Print("Thx for using the service")
			break
		}
	}
}

func sendMsgMenu() {
	scan := bufio.NewScanner(os.Stdin)
	var msg string
	for {
		fmt.Print("Write your massage: ")
		scan.Scan()
		msg = scan.Text()
		break
	}
	sendMsg(msg)
}

func sendMsg(msg string) {
	conn, err := net.DialTimeout("tcp", "localhost:6969", 3*time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

  //write message
	err = binary.Write(conn, binary.LittleEndian, uint32(len(msg)))
	if err != nil {
		panic(err)
	}
  
  err = conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
  if err != nil{
    panic(err)
  }
  
	_, err = conn.Write([]byte(msg))
	if err != nil {
		panic(err)
	}
  // read replay
	var size uint32
	err = binary.Read(conn, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}
  
  err = conn.SetReadDeadline(time.Now().Add(3 * time.Second))
  if err != nil{
    panic(err)
  }
  
	replay := make([]byte, size)
	_, err = conn.Read(replay)
	if err != nil {
		panic(err)
	}
	fmt.Println("Replay: " + string(replay))

}

func main() {
	menu()
}
