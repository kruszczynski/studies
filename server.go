package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/textproto"

	"github.com/kruszczynski/studies/stats"
)

type command int

const (
	rock command = iota
	paper
	scissors
	getStats
	quit
)

var commandNames = [5]string{
	"ROCK",
	"PAPER",
	"SCISSORS",
	"STATS",
	"QUIT",
}

var errInvalidCommand = errors.New("INV")

func main() {
	server()
}

func server() {
	ln, err := net.Listen("tcp", ":1983")
	if err != nil {
		fmt.Println(err)
		return
	}
	collector := stats.NewCollector()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleServerConnection(conn, collector)
	}
}

func handleServerConnection(conn net.Conn, collector *stats.Collector) {
	// greet the player
	conn.Write([]byte("Welcome stranger\n"))
	// receive the message
	r := bufio.NewReader(conn)
	textReader := textproto.NewReader(r)
	for {
		usedCommand, err := extractCommand(textReader)
		if err != nil {
			if err == errInvalidCommand {
				conn.Write([]byte("INVALID COMMAND\n"))
				continue
			} else {
				conn.Write([]byte("SOMEWHAT FATAL\n"))
				break
			}
		}
		handleCommand(usedCommand, conn, collector)

	}
}

func extractCommand(reader *textproto.Reader) (command, error) {
	line, err := reader.ReadLine()
	if err != nil {
		return getStats, err
	}

	for index, commandName := range commandNames {
		if commandName == line {
			return command(index), nil
		}
	}
	return getStats, errInvalidCommand
}

func handleCommand(usedCommand command, conn net.Conn, collector *stats.Collector) {
	if int(usedCommand) < 3 {
		// Let's play
		response := command(rand.Intn(3))
		responseString := commandNames[int(response)]
		switch {
		// DRAW
		case usedCommand == response:
			responseString = "DRAW " + responseString
			collector.DrawsCounter.Channel <- 1
		// USER LOSES
		case (usedCommand == 0 && response == 1) ||
			(usedCommand == 1 && response == 2) ||
			(usedCommand == 2 && response == 0):
			responseString = "LOSE " + responseString
			collector.LossesCounter.Channel <- 1
		// USER WINS, THE DEFAULT
		default:
			responseString = "WIN " + responseString
			collector.WinsCounter.Channel <- 1
		}
		conn.Write([]byte(responseString + "\n"))
	} else if usedCommand == getStats {
		conn.Write([]byte(collector.PrintStats() + "\n"))
	} else {
		conn.Write([]byte("Until then stranger\n"))
		conn.Close()
	}
}
