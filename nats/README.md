# Publish / Subscribe

1. Run NATS server using the following command
`docker run -p 4222:4222 -p 8222:8222 nats:alpine3.10`

2. Run subscriber
    ```
    cd subscriber
    go run subscriber.go
    ```
    This command will start subscriber waiting for messages on channel 

3. Run publisher
    ```
    cd publisher
    go run publisher.go
    ```
    This command will start publisher and the message type on the console will be published to the channel