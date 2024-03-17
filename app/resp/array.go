package resp

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type RespArray struct {
	Len  int
	Args []*RespBulkString
}

func NewRespArray(strs []string) *RespArray {
	respArr := &RespArray{
		Len: len(strs),
	}

	for _, s := range strs {
		respArr.Args = append(respArr.Args, NewRespBulkString(len(s), s))
	}

	return respArr
}

func DecodeToRespArray(str string) (*RespArray, error) {
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

func (ra *RespArray) Encode() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("*%d\r\n", ra.Len))

	for _, arg := range ra.Args {
		sb.WriteString(arg.Encode())
	}

	return sb.String()
}

func getLen(s string) (int, error) {
	r, _ := regexp.Compile(`[0-9]+`)
	match := r.FindString(s)

	return strconv.Atoi(match)
}
