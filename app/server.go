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

type ServerChannel struct {
	Data []byte
	Conn net.Conn
}

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
	log.Println("Server role: ", conf.Replication.Role)

	store := store.NewDataStore()
	serverCh := make(chan ServerChannel)

	go tcpConsumer(store, serverCh)

	if conf.Replication.Role == "slave" {
		c := replication.ConnectToMaster(args.masterHost, args.masterPort)
		replication.Handshake(c, port)

		go tcpProducer(c.GetConn(), serverCh)
	}

	defer l.Close()

	for {
		log.Println("Waiting for connections...")
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln("Error accepting connection: ", err.Error())
		}

		log.Println("Connection accepted")

		go tcpProducer(conn, serverCh)
	}
}

func tcpConsumer(store store.DataStore, ch chan ServerChannel) {
	for c := range ch {
		log.Println("Data from channel", string(c.Data))
		msgs, err := resp.RespParse(c.Data)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		for _, msg := range msgs {
			input := &commands.HandlerInput{
				Cmd:   msg.Command,
				Args:  msg.Args,
				Store: store,
				Conn:  c.Conn,
			}
			responses, err := commands.Handle(input)
			if err != nil {
				log.Printf("Error executing %s: %s\n", msg.Command, err.Error())
				continue
			}

			for _, response := range responses {
				n, err := c.Conn.Write([]byte(response))
				if err != nil {
					log.Fatalln("Error sending data: ", err.Error())
				}
				log.Printf("sent %d bytes\n", n)
			}
		}
	}
}

func tcpProducer(conn net.Conn, ch chan ServerChannel) {
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatalln("(handleConn) Error reading data: ", err.Error())
		}

		ch <- ServerChannel{
			Data: buffer,
			Conn: conn,
		}
	}
}
