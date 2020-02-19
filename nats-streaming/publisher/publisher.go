package main

import (
	"fmt"
	nats "github.com/nats-io/nats.go"
	stan "github.com/nats-io/stan.go"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// cluster ID should be the same as provided while running the NATS streaming server. Default is "test-cluster"
	clusterID := "cluster1"
	// clientID should be unique
	clientID := "test-client1"
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
	defer conn.Close()

	// Publish up to 10.
	for i := 0; i < 10; i++ {
		err := conn.Publish(channelSubject, []byte(fmt.Sprintf("%s %d", "Hello", i)))
		if err != nil {
			return err
		}
	}

	return nil
}
