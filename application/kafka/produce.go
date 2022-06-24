package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	route2 "github.com/jdaniloc/Individual-FullCycle-01/application/route"
	"github.com/jdaniloc/Individual-FullCycle-01/infra/kafka"
)

func Produce(msg *cKafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()

	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()

	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}

	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
