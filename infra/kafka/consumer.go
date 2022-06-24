package kafka

import (
	"fmt"
	"log"
	"os"

	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	MsgChan chan *cKafka.Message
}

func NewKafkaConsumer(msgChan chan *cKafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}

func (k *KafkaConsumer) Consume() {
	configMap := &cKafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupID"),
	}

	c, err := cKafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error consuming kafka message" + err.Error())
	}

	topics := []string{os.Getenv("KafkaReadTopic")}
	c.SubscribeTopics(topics, nil)
	fmt.Println("Kafka consumer has been started")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}
