package client

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

type RdbClient interface {
	Ping() string
	ReplConf([]string) string
	Set(key string, value string) string
	Get(key string) string
	PSync(replId string, offset int) (string, string)
	GetConn() net.Conn
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

func (c *client) PSync(replId string, offset int) (string, string) {
	log.Println("Sending PSYNC")

	msg := resp.NewRespArray([]string{"psync", replId, fmt.Sprintf("%d", offset)}).Encode()

	c.write(msg)
	data := c.read()
	rdbFile := c.read()

	return string(data), string(rdbFile)
}

func (c *client) Set(key string, value string) string {
	log.Println("Sending SET")

	msg := resp.NewRespArray([]string{"set", key, value}).Encode()

	c.write(msg)
	data := c.read()

	return string(data)
}

func (c *client) Get(key string) string {
	log.Println("Sending GET")

	msg := resp.NewRespArray([]string{"get", key}).Encode()

	c.write(msg)
	data := c.read()

	return string(data)
}

func (c *client) GetConn() net.Conn {
	return c.tcpConn
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
		if !errors.Is(err, io.EOF) {
			log.Fatalln("(Client) Error reading data", err.Error())
		}
	}

	log.Printf("%d bytes received from master\n", n)

	return data
}
