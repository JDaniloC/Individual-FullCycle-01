package kafka

import (
	"log"
	"os"

	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// NewKafkaProducer creates a ready to go kafka.Producer instance
func NewKafkaProducer() *cKafka.Producer {
	configMap := &cKafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		// "security.protocol": os.Getenv("security.protocol"),
		// "sasl.mechanisms":   os.Getenv("sasl.mechanisms"),
		// "sasl.username":     os.Getenv("sasl.username"),
		// "sasl.password":     os.Getenv("sasl.password"),
	}
	p, err := cKafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

// Publish is simple function created to publish new message to kafka
func Publish(msg string, topic string, producer *cKafka.Producer) error {
	message := &cKafka.Message{
		TopicPartition: cKafka.TopicPartition{
			Topic:     &topic,
			Partition: cKafka.PartitionAny,
		},
		Value: []byte(msg),
	}
	err := producer.Produce(message, nil)
	if err != nil {
		return err
	}
	return nil
}
