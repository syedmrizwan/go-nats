# go-nats

1. Run NATS server using the following command
`docker run -p 4222:4222 -p 8222:8222 nats:alpine3.10`

2. Run subscriber and give subject
```
cd subscriber
go run main.go hello
```

3. Run publisher and specify subject and message
```
cd publisher
go run main.go hello world
```