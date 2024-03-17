package commands

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/app/conf"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

func PSync() string {
	str := fmt.Sprintf("FULLRESYNC %s %d", conf.Replication.Id, 0)
	return resp.NewRespString(len(str), str).Encode()
}
