# Request / Reply
When an application sends a request and expects a response from a receiver, create two channels one for the request and one for the response.

1. Run NATS server using the following command
`docker run -p 4222:4222 -p 8222:8222 nats:alpine3.10`

2. Run reply
    ```
    cd reply
    go run reply.go
    ```
    This command will start the replier which will wait for requests. Every request recieved will be responded by the replier

3. Run requester
    ```
    cd request
    go run request.go
    ```
    This command will start requester and the message type on the console will be send as a request. Response from the receiver will be displayed on the console