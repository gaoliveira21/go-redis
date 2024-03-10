package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"

	"github.com/codecrafters-io/redis-starter-go/app/commands"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/store"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
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

			fmt.Println("Error reading data: ", err.Error())
			os.Exit(1)
		}

		parser := resp.NewRespParser(buffer)
		msg, err := parser.RespParse()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		var response string

		switch msg.Command {
		case "echo":
			r := commands.Echo(msg.Args)
			response = resp.NewRespString(r)
		case "ping":
			r := commands.Ping()
			response = resp.NewRespString(r)
		case "set":
			input := &commands.SetIput{
				Key:   msg.Args[0],
				Value: msg.Args[1],
			}

			if len(msg.Args) >= 4 {
				input.Exp, err = strconv.Atoi(msg.Args[3])
				if err != nil {
					fmt.Println("Could not convert exp to type int: ", err.Error())
					continue
				}
			}

			commands.Set(store, input)
			response = resp.NewRespString("OK")
		case "get":
			v, f := commands.Get(store, msg.Args[0])
			bs := resp.NewRespBulkString(len(v), v)

			if f {
				response = bs.Get()
			} else {
				response = bs.GetNull()
			}
		default:
			response = resp.NewRespString("Command not found")
		}

		n, err := conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Error sending data: ", err.Error())
			os.Exit(1)
		}
		fmt.Printf("sent %d bytes\n", n)
	}
}
