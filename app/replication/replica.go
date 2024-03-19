package replication

import (
	"log"
	"net"

	"github.com/codecrafters-io/redis-starter-go/app/replication/client"
)

type Replica struct {
	Conn net.Conn
}

var replicas []*Replica

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
	rdbClient.ReplConf([]string{"capa eof capa psync2", port})
	rdbClient.PSync("?", -1)
}

func AddReplica(r *Replica) {
	replicas = append(replicas, r)
}

func Propagate(b []byte) {
	for _, r := range replicas {
		log.Printf("Replica: %s\n", r.Conn.RemoteAddr().String())
		n, err := r.Conn.Write(b)
		if err != nil {
			log.Println("Error propagating to replica ", err.Error())
		} else {
			log.Printf("%d bytes sent to replica\n", n)
		}
	}
}
