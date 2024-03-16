package commands

import (
	"errors"

	"github.com/codecrafters-io/redis-starter-go/app/conf"
)

func Info(t string) (map[string]string, error) {
	switch t {
	case "replication":
		r := make(map[string]string)

		r["role"] = conf.Replication.Role

		return r, nil
	default:
		return nil, errors.New("invalid argument received")
	}
}
