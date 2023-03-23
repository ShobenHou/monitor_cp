package kafka

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type AgentConf struct {
	// Define your agent configuration structure here.
}

var kafkaProducer *kafka.Producer

func init() {
	// Initialize the Kafka producer.
	// Initialize the Kafka producer.
	var err error
	kafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
}

func PublishConfigToKafka(agentID string, agentConf AgentConf) error {
	agentConfJSON, err := json.Marshal(agentConf)
	if err != nil {
		return err
	}

	kafkaTopic := "agent-config"
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
		Key:            []byte(agentID),
		Value:          agentConfJSON,
	}

	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	err = kafkaProducer.Produce(msg, deliveryChan)
	if err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	return nil
}
