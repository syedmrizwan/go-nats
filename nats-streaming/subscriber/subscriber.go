package main

import (
	"fmt"
	"log"
	"runtime"

	nats "github.com/nats-io/nats.go"
	stan "github.com/nats-io/stan.go"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	//cluster ID should be the same as provided while running the NATS streaming server. Default is "test-cluster"
	clusterID := "cluster1"
	// clientID should be unique
	clientID := "test-client"
	// Channel subject
	channelSubject := "subject"
	conn, err := stan.Connect(
		clusterID, 
		clientID,  
		stan.NatsURL(nats.DefaultURL),
	)
	if err != nil {
		return err
	}

	conn.Subscribe(channelSubject, func(msg *stan.Msg) {
		// Print the value and whether it was redelivered.
		fmt.Printf("message = %s seq = %d [redelivered = %v]\n", string(msg.Data), msg.Sequence, msg.Redelivered)
	}, stan.DurableName("i-will-remember"))

	runtime.Goexit()
	return nil
}
