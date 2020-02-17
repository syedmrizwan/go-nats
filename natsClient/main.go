package main

import (
	"fmt"

	nats "github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, _ := nc.ChanSubscribe("foo", ch)

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	msg := <-ch
	fmt.Printf("Received a message: %s\n", string(msg.Data))

	// Unsubscribe
	sub.Unsubscribe()

	// Drain
	sub.Drain()

	// Close() not needed if this is called.
	nc.Drain()

	// Close connection
	nc.Close()

	select {}

}
