package resp

import (
	"regexp"
	"strconv"
	"strings"
)

type BulkString struct {
	Len   int
	Value string
}

type ArrayMessage struct {
	Len  int
	Args []BulkString
}

func getLen(s string) (int, error) {
	r, _ := regexp.Compile(`[0-9]+`)
	match := r.FindString(s)

	return strconv.Atoi(match)
}

func parseArray(data []byte) (*ArrayMessage, error) {
	str := string(data)
	lines := strings.Split(str, "\r\n")

	arrMsg := ArrayMessage{}

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
			arrMsg.Args = append(arrMsg.Args, BulkString{Len: l})
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
