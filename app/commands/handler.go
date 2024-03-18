package commands

import (
	"fmt"
	"strconv"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/store"
)

func Handle(cmd string, args []string, s store.DataStore) ([]string, error) {
	var response []string

	switch cmd {
	case "echo":
		response = append(response, Echo(args))
	case "ping":
		response = append(response, Ping())
	case "replconf":
		response = append(response, ReplConf())
	case "psync":
		r, rdbFile := PSync()
		response = append(response, r)
		response = append(response, fmt.Sprintf("$%d\r\n%s", len(rdbFile), rdbFile))
	case "set":
		input := &SetIput{
			Key:   args[0],
			Value: args[1],
		}

		if len(args) >= 4 {
			var err error
			input.Exp, err = strconv.Atoi(args[3])
			if err != nil {
				return []string{}, err
			}
		}

		response = append(response, Set(s, input))
	case "get":
		response = append(response, Get(s, args[0]))
	case "info":
		r, err := Info(args[0])
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
