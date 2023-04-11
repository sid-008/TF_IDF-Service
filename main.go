package main

import (
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	request := make([]byte, 1024)
	n, err := conn.Read(request) // n ---> number of bytes read

	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("%d bytes read", n)

	response := []byte("hello from golang microservice!")

	_, err = conn.Write(response)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}
