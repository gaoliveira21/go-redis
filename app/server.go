package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/codecrafters-io/redis-starter-go/app/commands"
	"github.com/codecrafters-io/redis-starter-go/app/conf"
	"github.com/codecrafters-io/redis-starter-go/app/replication"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/store"
)

func main() {
	serverArgs := GetServerArgs()

	if serverArgs.masterHost != "" {
		conf.Replication.Role = "slave"
	} else {
		conf.Replication.Role = "master"
		conf.Replication.Id = "8371b4fb1155b71f4a04d3e1bc3e18c4a990aeeb"
	}

	startServer(serverArgs)
}

func startServer(args *ServerArgs) {
	port := fmt.Sprintf("%d", args.port)

	l, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalln("Failed to bind to port " + port)
	}

	log.Println("Server listening on port " + port)

	if conf.Replication.Role == "slave" {
		c := replication.ConnecToMaster(args.masterHost, args.masterPort)
		replication.Handshake(c, port)

		defer c.Close()
	}

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
