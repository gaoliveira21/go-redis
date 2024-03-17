package commands

import (
	"strings"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

func Echo(args []string) string {
	str := strings.Join(args, "")

	rs := resp.NewRespString(len(str), str)

	return rs.Encode()
}
