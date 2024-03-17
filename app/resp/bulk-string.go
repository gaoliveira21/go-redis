package resp

import "fmt"

type RespBulkString struct {
	Len   int
	Value string
}

func NewRespBulkString(l int, v string) *RespBulkString {
	return &RespBulkString{
		Len:   l,
		Value: v,
	}
}

func (bs *RespBulkString) Encode() string {
	v := bs.Value

	return fmt.Sprintf("$%d\r\n%s\r\n", len(v), v)
}

func (bs *RespBulkString) EncodeNull() string {
	return "$-1\r\n"
}
