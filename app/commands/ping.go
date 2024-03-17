package commands

import "github.com/codecrafters-io/redis-starter-go/app/resp"

func Ping() string {
	return resp.NewRespString(4, "PONG").Encode()
}
