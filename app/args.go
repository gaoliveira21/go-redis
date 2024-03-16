package main

import (
	"flag"
	"log"
	"strconv"
)

type ServerArgs struct {
	port       int
	masterHost string
	masterPort int
}

func GetServerArgs() *ServerArgs {
	port := flag.Int("port", 6379, "Server port")
	masterHost := flag.String("replicaof", "", "")
	flag.Parse()

	serverArgs := &ServerArgs{
		port:       *port,
		masterHost: *masterHost,
	}

	args := flag.Args()

	if len(args) > 0 {
		masterPort, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalln("Invalid master port received ", err.Error())
		}

		serverArgs.masterPort = masterPort
	}

	return serverArgs
}
