package commands

import (
	"time"

	"github.com/codecrafters-io/redis-starter-go/app/store"
)

func Get(s store.DataStore, key string) (string, bool) {
	v, f := s[key]

	if f && !v.ExpiresAt.IsZero() {
		isExpired := v.ExpiresAt.Before(time.Now())

		if isExpired {
			delete(s, key)
			return v.Value, false
		}
	}

	return v.Value, f
}
