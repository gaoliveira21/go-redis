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

	return rdbClient
}

func Handshake(rdbClient client.RdbClient, port string) {
	log.Println("Start handshaking...")

	rdbClient.Ping()
	rdbClient.ReplConf([]string{"listening-port", port})
	rdbClient.ReplConf([]string{"capa psync2", port})
}
