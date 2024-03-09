package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/codecrafters-io/redis-starter-go/app/commands"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
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
		p, err := parser.RespParse()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		var resp string

		switch p.Command {
		case "echo":
			resp = commands.Echo(p.Args)
		case "ping":
			resp = commands.Ping()
		default:
			resp = "Command not found"
		}

		n, err := conn.Write([]byte("+" + resp + "\r\n"))
		if err != nil {
			fmt.Println("Error sending data: ", err.Error())
			os.Exit(1)
		}
		fmt.Printf("sent %d bytes\n", n)
	}
}
