package main

import (
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
	stats
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
	wins := make(chan int)
	losses := make(chan int)
	draws := make(chan int)
	go statAgregator(wins)
	go statAgregator(losses)
	go statAgregator(draws)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleServerConnection(conn, wins, losses, draws)
	}
}

func handleServerConnection(conn net.Conn, wins chan int, losses chan int, draws chan int) {
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
			case usedCommand == response:
				responseString = "DRAW " + responseString
				draws <- 1
			case (usedCommand == 0 && response == 1) ||
				(usedCommand == 1 && response == 2) ||
				(usedCommand == 2 && response == 0):
				responseString = "LOSE " + responseString
				losses <- 1
			default:
				responseString = "WIN " + responseString
				wins <- 1
			}
			conn.Write([]byte(responseString + "\n"))
		} else {
			winsTotal := <-wins
			// lossesTotal := <-losses
			// drawsTotal := <-draws
			// response := fmt.Sprintf("W%d L%d D%d", winsTotal, lossesTotal, drawsTotal)
			fmt.Println(winsTotal)
			// conn.Write([]byte(response + "\n"))
		}

	}
}

func extractCommand(reader *bufio.Reader) (command, error) {
	line, err := reader.ReadBytes('\n')
	if err != nil {
		return stats, err
	}
	lineString := strings.Trim(string(line), "\n\r\t")

	for index, commandName := range commandNames {
		if commandName == lineString {
			return command(index), nil
		}
	}
	return stats, errors.New("INV")
}

func statAgregator(channel chan int) {
	counter := 0
	for {
		increment := <-channel
		counter = counter + increment
	}
}
