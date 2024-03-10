package store

import "time"

type Data struct {
	Value     string
	ExpiresAt time.Time
}

type DataStore map[string]Data

func NewDataStore() DataStore {
	return make(DataStore)
}
