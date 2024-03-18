package commands

import (
	"encoding/hex"
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/app/conf"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

var EMPTY_RDB_FILE, _ = hex.DecodeString("524544495330303131fa0972656469732d76657205372e322e30fa0a72656469732d62697473c040fa056374696d65c26d08bc65fa08757365642d6d656dc2b0c41000fa08616f662d62617365c000fff06e3bfec0ff5aa2")

func PSync() (string, []byte) {
	str := fmt.Sprintf("FULLRESYNC %s %d", conf.Replication.Id, 0)
	return resp.NewRespString(len(str), str).Encode(), EMPTY_RDB_FILE
}
