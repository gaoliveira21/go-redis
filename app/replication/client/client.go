package client

import (
	"fmt"
	"log"
	"net"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

type RdbClient interface {
	Ping() string
	Close() error
}

type client struct {
	tcpConn net.Conn
}

func (c *client) Close() error {
	return c.tcpConn.Close()
}

func (c *client) Ping() string {
	msg := resp.NewRespArray([]string{"ping"}).Encode()

	n, err := c.tcpConn.Write([]byte(msg))
	if err != nil {
		log.Println("Ping: command fail ", err.Error())
		return ""
	}

	log.Printf("Ping: %d bytes sent to master\n", n)

	data := make([]byte, 1024)
	n, err = c.tcpConn.Read(data)
	if err != nil {
		log.Println("Ping: command fail ", err.Error())
		return ""
	}

	log.Printf("Ping: %d bytes received from master\n", n)

	return string(data)
}

func Connect(host string, port int) (RdbClient, error) {
	tcp, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}

	return &client{
		tcpConn: tcp,
	}, nil
}
