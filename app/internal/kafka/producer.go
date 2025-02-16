package kafka

import (
	"fmt"
)

type Message struct {
	Topic     string
	Partition int
	Key       string
	Value     string
}

type ProducerRepositoryInterface interface {
	ProduceMessage(jobID int, key string, action string) error
	// FetchJobs(topic string, id int) error
}

type NewProducer struct {
	DefaultTopic     string
	DefaultPartition int
}

func (p *NewProducer) produce(msg Message) error {
	fmt.Printf(
		"Producing message to topic '%s' [Partition: %d] - Key: %s, Value: %s\n",
		msg.Topic, msg.Partition, msg.Key, msg.Value,
	)
	return nil
}

func NewKafkaProducer(topic string, partition int) *NewProducer {
	// kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
	// 	"bootstrap.servers": "localhost:9092"})
	// if err != nil {
	// 	panic(err)
	// }
	return &NewProducer{
		DefaultTopic:     topic,
		DefaultPartition: partition,
	}
}

func (p *NewProducer) ProduceMessage(jobID int, key string, action string) error {
	message := Message{
		Topic:     p.DefaultTopic,
		Partition: p.DefaultPartition,
		Key:       key,
		Value:     action,
	}
	return p.produce(message)
}
