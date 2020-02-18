package main

import (
	"log"

	stan "github.com/nats-io/stan.go"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	conn, err := stan.Connect(
		"cluster1",//cluster ID should be the same as provided while running the NATS streaming server. Default is "test-cluster"
		"test-client1",// need to give unique ID
		stan.NatsURL("nats://localhost:4222"),
	)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Publish up to 10.
	for i := 0; i < 10; i++ {
		err := conn.Publish("counter-channel", nil)
		if err != nil {
			return err
		}
	}

	// Wait until all messages have been processed.
	//wg.Wait()

	return nil
}
