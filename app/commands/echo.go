package commands

import (
	"strings"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

func Echo(args []string) string {
	str := strings.Join(args, "")

	return resp.NewRespString(len(str), str).Encode()
}
