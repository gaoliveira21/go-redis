package resp

import "fmt"

type RespBulkString struct {
	Len   int
	Value string
}

func NewRespBulkString(l int, v string) RespBulkString {
	return RespBulkString{
		Len:   l,
		Value: v,
	}
}

func (bs *RespBulkString) Get() string {
	v := bs.Value

	return "$" + fmt.Sprintf("%d", len(v)) + "\r\n" + v + "\r\n"
}

func (bs *RespBulkString) GetNull() string {
	return "$-1\r\n"
}
