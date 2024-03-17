package commands

import "github.com/codecrafters-io/redis-starter-go/app/resp"

func ReplConf() string {
	return resp.NewRespString(2, "OK").Encode()
}
