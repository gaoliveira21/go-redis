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
		rs := resp.NewRespString(len(r), r)
		response = rs.Encode()
	case "ping":
		r := Ping()
		rs := resp.NewRespString(len(r), r)
		response = rs.Encode()
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
		rs := resp.NewRespString(2, "OK")
		response = rs.Encode()
	case "get":
		v, f := Get(s, args[0])
		bs := resp.NewRespBulkString(len(v), v)

		if f {
			response = bs.Encode()
		} else {
			response = bs.EncodeNull()
		}
	default:
		msg := "Command not found"
		rs := resp.NewRespString(len(msg), msg)
		response = rs.Encode()
	}

	return response, nil
}
