package main

import (
	"../config"
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type cfile = config.Config

func readF(s string) []cfile {
	return config.ReadFile(s)
}

//function that prints the message on the server side
func Messageprint(ID int, Message string) {

	fmt.Println("\n ---------------------------- \n --- New Incoming Message --- \n ---------------------------- \n")
	fmt.Println("Message from:" + strconv.Itoa(ID))
	fmt.Println("Message Content: " + Message)
	t := time.Now()
	myTime := t.Format(time.RFC850) + "\n"
	fmt.Println("Confirmed received at: " + myTime)
	fmt.Println("\n -------------------------- \n")

}

//takes the message and then sends it to be printed
func UnicastReceive(config cfile, message string) {
	var identifier int
	identifier = config.ID
	Messageprint(identifier, message)
}

func main() {
	var config cfile

	//read info from config file
	config = readF("config.txt")[0]

	//create TCP listener
	address := config.IP + ":" + config.Port
	ln, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		//listen for communication
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		for {
			//read the incoming information with protocol in case of error
			netData, err := bufio.NewReader(conn).ReadString('\n')

			if err != nil {
				fmt.Println(err)
				return
			}

			//termination protocol allows the client to end connection manually
			if strings.TrimSpace(netData) != "END" {
				//if termination protocol is not called then receive and print the message
				UnicastReceive(config, netData)
			} else if strings.TrimSpace(netData) == "END"{
				fmt.Println("Exiting TCP server!")
				return
			}
		}
	}
}
