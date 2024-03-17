package commands

import (
	"time"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/store"
)

func Get(s store.DataStore, key string) string {
	v, f := s[key]

	bs := resp.NewRespBulkString(len(v.Value), v.Value)

	if f && !v.ExpiresAt.IsZero() {
		isExpired := v.ExpiresAt.Before(time.Now())

		if isExpired {
			delete(s, key)
			return bs.EncodeNull()
		}
	}

	if f {
		return bs.Encode()
	}

	return bs.EncodeNull()
}
