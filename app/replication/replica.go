package replication

import (
	"fmt"
	"log"

	"github.com/codecrafters-io/redis-starter-go/app/replication/client"
)

func ConnecToMaster(host string, port int) client.RdbClient {
	log.Printf("Connecting to master at %s:%d\n", host, port)

	rdbClient, err := client.Connect(host, port)
	if err != nil {
		log.Fatalf("Failed to connect to master at %s:%d\n", host, port)
	}

	Handshake(rdbClient, port)

	return rdbClient
}

func Handshake(rdbClient client.RdbClient, port int) {
	log.Println("Start handshaking...")

	rdbClient.Ping()
	rdbClient.ReplConf([]string{"listening-port", fmt.Sprintf("%d", port)})
	rdbClient.ReplConf([]string{"capa psync2", fmt.Sprintf("%d", port)})
}
