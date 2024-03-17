package commands

import "github.com/codecrafters-io/redis-starter-go/app/resp"

func Ping() string {
	rs := resp.NewRespString(4, "PONG")
	return rs.Encode()
}
