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
		r := Echo(args)
		response = resp.NewRespString(r)
	case "ping":
		r := Ping()
		response = resp.NewRespString(r)
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

		Set(s, input)
		response = resp.NewRespString("OK")
	case "get":
		v, f := Get(s, args[0])
		bs := resp.NewRespBulkString(len(v), v)

		if f {
			response = bs.Get()
		} else {
			response = bs.GetNull()
		}
	default:
		response = resp.NewRespString("Command not found")
	}

	return response, nil
}
