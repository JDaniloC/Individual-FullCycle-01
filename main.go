package main

import (
	"fmt"
	"log"

	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	kafka2 "github.com/jdaniloc/Individual-FullCycle-01/application/kafka"
	"github.com/jdaniloc/Individual-FullCycle-01/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	msgChan := make(chan *cKafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}
