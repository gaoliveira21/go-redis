package client

import (
	"fmt"
	"log"
	"net"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

type RdbClient interface {
	Ping() string
	ReplConf([]string) string
	Close() error
}

type client struct {
	tcpConn net.Conn
}

func (c *client) Ping() string {
	log.Println("Sending PING")

	msg := resp.NewRespArray([]string{"ping"}).Encode()

	c.write(msg)
	data := c.read()

	return string(data)
}

func (c *client) ReplConf(s []string) string {
	log.Println("Sending REPLCONF")

	msg := resp.NewRespArray(append([]string{"replconf"}, s...)).Encode()

	c.write(msg)
	data := c.read()

	return string(data)
}

func (c *client) Close() error {
	return c.tcpConn.Close()
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

func (c *client) write(msg string) {
	n, err := c.tcpConn.Write([]byte(msg))
	if err != nil {
		log.Fatalln("Error writing data", err.Error())
	}

	log.Printf("%d bytes sent to master\n", n)
}

func (c *client) read() []byte {
	data := make([]byte, 1024)
	n, err := c.tcpConn.Read(data)
	if err != nil {
		log.Fatalln("Error reading data", err.Error())
	}

	log.Printf("%d bytes received from master\n", n)

	return data
}
