package replication

import (
	"log"

	"github.com/codecrafters-io/redis-starter-go/app/replication/client"
)

func ConnecToMaster(host string, port int) client.RdbClient {
	log.Printf("Connecting to master at %s:%d\n", host, port)

	rdbClient, err := client.Connect(host, port)
	if err != nil {
		log.Fatalf("Failed to connect to master at %s:%d\n", host, port)
	}

	Handshake(rdbClient)

	return rdbClient
}

func Handshake(rdbClient client.RdbClient) {
	log.Println("Start handshaking...")

	ping := rdbClient.Ping()
	log.Printf("PING: %s\n", ping)
}
