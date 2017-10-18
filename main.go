package main

import (
	"github.com/name5566/leaf/log"
	"os"
	"os/signal"
	"leafclient/node"
)

const version = "0.0.1"

func main() {
	//logger
	log.Release("Client %v starting up", version)

	//clients online
	node.Online()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Release("Client closing down (signal: %v)", sig)
	node.Destroy()
}
