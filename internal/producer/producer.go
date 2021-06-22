package producer

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
	"time"
)

type EventType int

const (
	Created EventType = iota
	Updated
	Removed
	Described
)

type Message struct {
	Type      EventType
	Id        uint64
	Timestamp uint64
}

type Producer interface {
	SendMessage(msg Message) bool
	Close() error
	CreateMessage(Type EventType, id uint64) Message
}

type producer struct {
	kafkaProd sarama.SyncProducer
	topic     string
}

func NewProducer(brokers []string, topic string) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	kafkaProd, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &producer{
		kafkaProd: kafkaProd,
		topic:     topic,
	}, nil
}

func (prod *producer) SendMessage(msg Message) bool {
	json, err := json.Marshal(msg)

	if err != nil {
		return false
	}

	saramaMsg := sarama.ProducerMessage{
		Topic:     prod.topic,
		Partition: -1,
		Value:     sarama.StringEncoder(json),
	}

	_, _, err = prod.kafkaProd.SendMessage(&saramaMsg)
	if err != nil {
		log.Printf(err.Error())
		return false
	}

	return true
}

func (prod *producer) Close() error {
	return prod.kafkaProd.Close()
}

func (prod *producer) CreateMessage(Type EventType, id uint64) Message {
	return Message{
		Type:      Type,
		Id:        id,
		Timestamp: uint64(time.Now().UTC().Unix()),
	}
}
