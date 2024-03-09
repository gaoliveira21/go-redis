package resp

import (
	"regexp"
	"strconv"
	"strings"
)

type RespArray struct {
	Len  int
	Args []RespBulkString
}

func getLen(s string) (int, error) {
	r, _ := regexp.Compile(`[0-9]+`)
	match := r.FindString(s)

	return strconv.Atoi(match)
}

func parseArray(data []byte) (*RespArray, error) {
	str := string(data)
	lines := strings.Split(str, "\r\n")

	arrMsg := RespArray{}

	for _, line := range lines {
		if line[0] == '*' {
			l, err := getLen(line)
			if err != nil {
				return nil, err
			}
			arrMsg.Len = l
			continue
		}

		if line[0] == '$' {
			l, err := getLen(line)
			if err != nil {
				return nil, err
			}
			arrMsg.Args = append(arrMsg.Args, NewRespBulkString(l, ""))
			continue
		}

		i := len(arrMsg.Args) - 1
		arg := arrMsg.Args[i]
		if arg.Len == len(line) {
			arg.Value = strings.ToLower(line)
			arrMsg.Args[i] = arg
		}
	}

	return &arrMsg, nil
}
