package commands

import (
	"strconv"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/store"
)

func Handle(cmd string, args []string, s store.DataStore) (string, error) {
	var response string

	switch cmd {
	case "echo":
		response = Echo(args)
	case "ping":
		response = Ping()
	case "replconf":
		response = ReplConf()
	case "psync":
		response = PSync()
	case "set":
		input := &SetIput{
			Key:   args[0],
			Value: args[1],
		}

		if len(args) >= 4 {
			var err error
			input.Exp, err = strconv.Atoi(args[3])
			if err != nil {
				return "", err
			}
		}

		response = Set(s, input)
	case "get":
		response = Get(s, args[0])
	case "info":
		r, err := Info(args[0])
		if err != nil {
			return "", err
		}

		response = r
	default:
		msg := "Command not found"
		response = resp.NewRespString(len(msg), msg).Encode()
	}

	return response, nil
}
