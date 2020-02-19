package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	nats "github.com/nats-io/nats.go"
)

func main() {

	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Sample Publisher")}

	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Channel subject
	channelSubject := "subject"

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("NATS REQUEST")
	fmt.Println("---------------------")
	for {
		fmt.Print("-> ")
		payload, _ := reader.ReadString('\n')
		// convert CRLF to LF
		payload = strings.Replace(payload, "\n", "", -1)

		msg, err := nc.Request(channelSubject, []byte(payload), 2*time.Second)
		if err != nil {
			if nc.LastError() != nil {
				log.Fatalf("%v for request", nc.LastError())
			}
			log.Fatalf("%v for request", err)
		}
		nc.Flush()

		if err := nc.LastError(); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Published [%s] : '%s'", channelSubject, payload)
			log.Printf("Received  [%v] : '%s'", msg.Subject, string(msg.Data))
		}
	}
}
