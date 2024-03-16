package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/codecrafters-io/redis-starter-go/app/commands"
	"github.com/codecrafters-io/redis-starter-go/app/conf"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/store"
)

func main() {
	serverArgs := GetServerArgs()
	conf.Replication = &conf.ReplicationConf{
		Role: "master",
	}

	if serverArgs.masterHost != "" {
		conf.Replication.Role = "slave"
	}

	strPort := fmt.Sprintf("%d", serverArgs.port)

	l, err := net.Listen("tcp", "0.0.0.0:"+strPort)
	if err != nil {
		log.Fatalln("Failed to bind to port " + strPort)
	}

	log.Println("Server listening on port " + strPort)

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
