package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	server()
}

func server() {
	ln, err := net.Listen("tcp", ":1983")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleServerConnection(conn)
	}
}

func handleServerConnection(conn net.Conn) {
	// greet the player
	conn.Write([]byte("Welcome stranger\n"))
	// receive the message
	b := bufio.NewReader(conn)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil { // EOF, or worse
			break
		}
		conn.Write(line)
	}
}
