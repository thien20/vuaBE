package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type ProducerRepositoryInterface interface {
	// ProduceMessage(topic string, message []byte) error
	FetchJobs(topic string, id int) error
}

type producer struct {
	producer *kafka.Producer
	topic    string
}

func NewProducer(topic string) (*producer, error) {
	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
	return &producer{producer: kafkaProducer, topic: topic}, nil
}

func (p *producer) ProduceMessage(jobID int, key string, action string) error {
	message := fmt.Sprintf(`{"jobID": %d, "action": "%s"}`, jobID, action)
	p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(message),
	}, nil)

	p.producer.Flush(1 * 1000)
	return nil
}
