package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/app/conf"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

func Info(t string) (string, error) {
	switch t {
	case "replication":
		var sb strings.Builder

		sb.WriteString("# Replication \n")
		sb.WriteString(fmt.Sprintf("role:%s\n", conf.Replication.Role))
		sb.WriteString(fmt.Sprintf("master_replid:%s\n", conf.Replication.Id))
		sb.WriteString(fmt.Sprintf("master_repl_offset:%v", conf.Replication.Offset))

		r := resp.NewRespBulkString(len(sb.String()), sb.String()).Encode()

		return r, nil
	default:
		return "", errors.New("invalid argument received")
	}
}
