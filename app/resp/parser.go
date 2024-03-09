package resp

import (
	"errors"
)

type RESP struct {
	data []byte
}

type RespMessage struct {
	Command string
	Args    []string
}

func NewRespParser(b []byte) *RESP {
	return &RESP{
		data: b,
	}
}

func (r *RESP) RespParse() (*RespMessage, error) {
	switch r.data[0] {
	case '*':
		{
			r, err := parseArray(r.data)
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
