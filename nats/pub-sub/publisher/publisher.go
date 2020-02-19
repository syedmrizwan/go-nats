package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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
	channelSubject := "channelsubject"

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("NATS PUBLISHER")
	fmt.Println("---------------------")
	for {
		fmt.Print("-> ")
		msg, _ := reader.ReadString('\n')
		// convert CRLF to LF
		msg = strings.Replace(msg, "\n", "", -1)

		nc.Publish(channelSubject, []byte(msg))
		nc.Flush()

		if err := nc.LastError(); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Published [%s] : '%s'\n", channelSubject, msg)
		}
	}
}
