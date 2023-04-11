package main

// TODO Better connection to microservice

import (
	"log"
	"net"

	"github.com/sid-008/CN-project/service"
)

func handleConnection(conn net.Conn) {
	request := make([]byte, 1024)
	n, err := conn.Read(request) // n ---> number of bytes read

	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("%d bytes read", n)
	//log.Println(string(request))
	log.Printf("%T", request)

	service.InsertToDB(request)

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
	log.Println("Server started on localhost 8000, listening for incoming requests")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}
