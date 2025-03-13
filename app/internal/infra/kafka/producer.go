package kafka

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
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
	Broker           string
	DefaultTopic     string
	DefaultPartition int
	Producer         *kafka.Producer
}

func NewKafkaProducer(broker string, topic string, partition int) *NewProducer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		panic(err)
	}
	return &NewProducer{
		Broker:           broker,
		DefaultTopic:     topic,
		DefaultPartition: partition,
		Producer:         p,
	}
}

func (p *NewProducer) produce(msg Message) error {
	log.Printf("Producing message: topic=%s, key=%s, value=%s", msg.Topic, msg.Key, msg.Value)

	kafkaMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &msg.Topic, Partition: int32(msg.Partition)},
		Key:            []byte(msg.Key),
		Value:          []byte(msg.Value),
	}
	err := p.Producer.Produce(kafkaMsg, nil)
	if err != nil {
		return err
	}
	// Flush to ensure delivery
	// p.Producer.Flush(15 * 1000)
	return nil
}

func (p *NewProducer) ProduceMessage(jobID int, key string, action string) error {
	messageData, err := json.Marshal(
		map[string]interface{}{
			"action": action,
			"job_id": jobID})
	if err != nil {
		return err
	}

	message := Message{
		Topic:     p.DefaultTopic,
		Partition: p.DefaultPartition,
		Key:       key,
		Value:     string(messageData),
	}
	return p.produce(message)
}
