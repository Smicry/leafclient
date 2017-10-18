package main

import (
	"os"
	"os/signal"
	"github.com/name5566/leaf/log"
)

const version = "0.0.1"

func main() {
	//logger
	log.Release("Node %v starting up", version)

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Release("Node closing down (signal: %v)", sig)
	//Destroy()
}
