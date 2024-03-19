package resp

import (
	"errors"
)

type RespMessage struct {
	Command string
	Args    []string
}

func RespParse(data []byte) ([]*RespMessage, error) {
	switch data[0] {
	case '*':
		{
			r, err := DecodeToRespArray(string(data))
			if err != nil {
				return nil, err
			}

			var msgs []*RespMessage
			for _, v := range r {
				args := []string{}

				for _, v := range v.Args[1:] {
					args = append(args, v.Value)
				}

				msgs = append(msgs, &RespMessage{Command: v.Args[0].Value, Args: args})
			}

			return msgs, nil
		}
	default:
		return nil, errors.New("RespParse: data could not be parsed")
	}
}
