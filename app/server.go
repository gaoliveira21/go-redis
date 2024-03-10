package main

import (
	"errors"
	"io"
	"log"
	"net"

	"github.com/codecrafters-io/redis-starter-go/app/commands"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/store"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		log.Fatalln("Failed to bind to port 6379")
	}

	log.Println("Server listening on port 6379")

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln("Error accepting connection: ", err.Error())
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	store := store.NewDataStore()

	for {
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatalln("Error reading data: ", err.Error())
		}

		msg, err := resp.RespParse(buffer)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		response, err := commands.Handle(msg.Command, msg.Args, store)
		if err != nil {
			log.Printf("Error executing %s: %s\n", msg.Command, err.Error())
			continue
		}

		n, err := conn.Write([]byte(response))
		if err != nil {
			log.Fatalln("Error sending data: ", err.Error())
		}
		log.Printf("sent %d bytes\n", n)
	}
}
