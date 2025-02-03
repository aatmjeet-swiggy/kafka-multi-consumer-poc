# Kafka multiconsumer POC

This repository demonstrates a proof of concept for a Kafka multi-consumer setup.

## Installation

To install the necessary Go modules, run:
```
go mod tidy
```

To setup depenedencies (kafka and zoo-keeper) run:
```
podman-compose up -d

or 
docker-compose up -d
```

To run 2 producers run:
```
1st producer:
go run producer/basic/basic_producer.go
```

```
2nd producer
go run producer/basic_replica/basic_producer.go
```

Finally to check consumers:

To run `swgykafka`'s CG, run:
```
go run consumer/cosumer.go
```

To run `sarama`'s CG, run:
```
go run consumer/sarama-consumer/consumer.go
```