package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	nats "github.com/nats-io/nats.go"
	stan "github.com/nats-io/stan.go"
)

func main() {
	var client string
	flag.StringVar(&client, "client", "test-client-1", "Nats client Id")
	flag.Parse()

	if err := run(client); err != nil {
		log.Fatal(err)
	}
}

func run(clientID string) error {
	//cluster ID should be the same as provided while running the NATS streaming server. Default is "test-cluster"
	clusterID := "cluster1"

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

	conn.QueueSubscribe(channelSubject, "Queue1", func(msg *stan.Msg) {
		// Print the value and whether it was redelivered.
		fmt.Printf("message = %s seq = %d [redelivered = %v]\n", string(msg.Data), msg.Sequence, msg.Redelivered)
		time.Sleep(10 * time.Second)
		msg.Ack()
	}, stan.SetManualAckMode(), stan.AckWait(stan.DefaultAckWait), stan.DurableName("i-will-remember"))

	runtime.Goexit()
	return nil
}
