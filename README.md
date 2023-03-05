# FullCycle training
Learning Go, Apache Kafka and Kubernets

## Simulator and Kafka working

Explain how the simulator works and how to run the kafka cluster. Where the kafka is responsible for the data travel between the applications, and the simulator will receive a route through topic `new-direction` and send coordinates to the topic `new-position`.

### To run the simulator independently

Run the Apacha Kafka with the following command:
```bash
cd apache-kafka
docker-compose up -d
```

Enter in the Kafka container and run the producer:
```bash
docker exec -it apache-kafka_kafka_1 bash
kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction
```

Open another terminal and run the simulator in the main folder:
```bash
docker-compose up -d
docker exec -it simulator /bin/bash
go run main.go
```

With the kafka cluster running, open another terminal and run the consumer:
```bash
docker exec -it apache-kafka_kafka_1 bash
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position
```

### How the simulator works

1. A producer sends a message with the following format to the `route.new-direction` topic:
    ```json
    {"clientId": "1", "routeId": "1"}
    ```
    Where `clientId` is the id of the client that sends the message and `routeId` is the id of the [route](./destinations/).

2. The simulator will read the JSON in the `./destinations/{id}.txt`
3. For each coordinate it will send a message with the following format to the `route.new-position` topic:
    ```json
    {"clientId": "1", "routeId": "1", "position": [0, 0], "finished": false}
    ```
    When the "finished" is true it's because the route ends.
