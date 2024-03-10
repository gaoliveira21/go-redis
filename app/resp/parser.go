package resp

import (
	"errors"
)

type RespMessage struct {
	Command string
	Args    []string
}

func RespParse(data []byte) (*RespMessage, error) {
	switch data[0] {
	case '*':
		{
			r, err := parseArray(data)
			if err != nil {
				return nil, err
			}

			args := []string{}

			for _, v := range r.Args[1:] {
				args = append(args, v.Value)
			}

			return &RespMessage{Command: r.Args[0].Value, Args: args}, nil
		}
	default:
		return nil, errors.New("RespParse: data could not be parsed")
	}
}
