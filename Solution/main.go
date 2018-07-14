package main

import (
	"fmt"
	"log"
	"net"
)
const (
	// Connection parameters
	PORT = "8080"
	HOST = "localhost"
	TYPE = "tcp"
)
func main() {
	//server code
	ln, err := net.Listen(TYPE, HOST+":"+PORT)
	defer func() {
		ln.Close()
	}()
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Printf("Server is listening on port %s...", PORT)
	ManageClient(ln)
}


// ManageClient  handles client connections.

func ManageClient(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			break
		}
		ClientController(conn)
	}
}
