package resp

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type RESP struct {
	data []byte
}

type ParseOutput struct {
	Command string
	Args    []string
}

func NewRespParser(b []byte) *RESP {
	return &RESP{
		data: b,
	}
}

func (r *RESP) RespParse() (*ParseOutput, error) {
	switch r.data[0] {
	case '*':
		r, err := parseArray(r.data)
		if err != nil {
			return nil, err
		}
		return &ParseOutput{Command: r[0], Args: r[1:]}, nil
	}

	return nil, errors.New("data could not be parsed")
}

func parseArray(data []byte) ([]string, error) {
	str := string(data)
	lines := strings.Split(str, "\\r\\n")

	l, err := getArrayLen(lines[0])
	if err != nil {
		return []string{}, err
	}

	out := []string{}
	args := lines[1:]

	for i := 0; i < l*2; i++ {
		if i >= len(args) {
			return []string{}, errors.New("data could not be parsed")
		}

		arg := args[i]
		if arg[0] != '$' {
			out = append(out, arg)
		}
	}

	return out, nil
}

func getArrayLen(s string) (int, error) {
	r, _ := regexp.Compile(`[0-9]+`)
	match := r.FindString(s)

	return strconv.Atoi(match)
}
