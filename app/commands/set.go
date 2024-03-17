package commands

import (
	"time"

	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/store"
)

type SetIput struct {
	Key   string
	Value string
	Exp   int
}

func Set(s store.DataStore, i *SetIput) string {
	var exp time.Time

	if i.Exp != 0 {
		exp = time.Now().Add(time.Duration(i.Exp) * time.Millisecond)
	}

	s[i.Key] = store.Data{
		Value:     i.Value,
		ExpiresAt: exp,
	}

	return resp.NewRespString(2, "OK").Encode()
}
