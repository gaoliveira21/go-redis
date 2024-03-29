package main

import (
	"log"
	"testing"
	"time"

	"github.com/codecrafters-io/redis-starter-go/app/replication/client"
)

func TestReplication(t *testing.T) {
	mc, err := client.Connect("localhost", 6379)

	if err != nil {
		t.Error("Failed to connect to master")
	}

	sc1, _ := client.Connect("localhost", 6380)
	sc2, _ := client.Connect("localhost", 6381)

	mc.Set("foo", "123")
	mc.Set("bar", "456")
	mc.Set("baz", "789")

	time.Sleep(3000)

	r := mc.Get("foo")
	r1 := sc1.Get("foo")
	r2 := sc2.Get("foo")

	log.Println(r, r1, r2)

	r = mc.Get("bar")
	r1 = sc1.Get("bar")
	r2 = sc2.Get("bar")

	log.Println(r, r1, r2)

	r = mc.Get("baz")
	r1 = sc1.Get("baz")
	r2 = sc2.Get("baz")

	log.Println(r, r1, r2)
}
