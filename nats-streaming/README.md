# Durable Subscriber with NATS Streaming

With this example, a subscriber is disconnected after 5 messages, reconnects and processes the remaining messages.

1. Run NATS streaming server using the following command. 

    `docker run -p 4222:4222 -p 8222:8222 nats-streaming -store file -m 8222 -dir datastore --cluster_id cluster1`

    Name used after `--cluster_id` flag will be needed to establish a connection. Default cluster_id is `test-cluster`

2. Run listener
    ```
    cd listener
    go run listener.go
    ```

3. Run publisher 
    ```
    cd publisher
    go run publisher.go 
    ```