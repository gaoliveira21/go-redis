package commands

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/codecrafters-io/redis-starter-go/app/conf"
	"github.com/codecrafters-io/redis-starter-go/app/replication"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/store"
)

type HandlerInput struct {
	Cmd   string
	Args  []string
	Store store.DataStore
	Conn  net.Conn
}

func Handle(i *HandlerInput) ([]string, error) {
	var response []string

	log.Printf("Command %s received\n", i.Cmd)

	switch i.Cmd {
	case "echo":
		response = append(response, Echo(i.Args))
	case "ping":
		response = append(response, Ping())
	case "replconf":
		response = append(response, ReplConf())
	case "psync":
		r, rdbFile := PSync(i.Conn)
		response = append(response, r)
		response = append(response, fmt.Sprintf("$%d\r\n%s", len(rdbFile), rdbFile))
	case "set":
		input := &SetIput{
			Key:   i.Args[0],
			Value: i.Args[1],
		}

		if len(i.Args) >= 4 {
			var err error
			input.Exp, err = strconv.Atoi(i.Args[3])
			if err != nil {
				return []string{}, err
			}
		}

		r := Set(i.Store, input)
		if conf.Replication.Role == "master" {
			response = append(response, r)
			ra := resp.NewRespArray(append([]string{i.Cmd}, i.Args...)).Encode()
			replication.Propagate([]byte(ra))
		}
	case "get":
		response = append(response, Get(i.Store, i.Args[0]))
	case "info":
		r, err := Info(i.Args[0])
		if err != nil {
			return []string{}, err
		}

		response = append(response, r)
	default:
		msg := "Command not found"
		response = append(response, resp.NewRespString(len(msg), msg).Encode())
	}

	return response, nil
}
