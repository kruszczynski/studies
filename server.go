package main

import (
	"./stats"
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"strings"
)

type command int

const (
	rock command = iota
	paper
	scissors
	getStats
)

var commandNames = [4]string{
	"ROCK",
	"PAPER",
	"SCISSORS",
	"STATS",
}

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
	for {
		usedCommand, err := extractCommand(r)
		if err != nil {
			if err.Error() == "INV" {
				conn.Write([]byte("INVALID COMMAND\n"))
				continue
			} else {
				conn.Write([]byte("SOMEWHAT FATAL\n"))
				break
			}
		}
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
		} else {
			conn.Write([]byte(collector.PrintStats() + "\n"))
		}

	}
}

func extractCommand(reader *bufio.Reader) (command, error) {
	line, err := reader.ReadBytes('\n')
	if err != nil {
		return getStats, err
	}
	lineString := strings.Trim(string(line), "\n\r\t")

	for index, commandName := range commandNames {
		if commandName == lineString {
			return command(index), nil
		}
	}
	return getStats, errors.New("INV")
}
