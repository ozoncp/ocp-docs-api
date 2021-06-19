package producer

import (
"github.com/Shopify/sarama"
"github.com/rs/zerolog/log"
)

type Producer interface {
	SendMessage(msg string) bool
	Close() error
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

func (prod *producer) SendMessage(msg string) bool {
	saramaMsg := prepareMessage(prod.topic, msg)
	_, _, err := prod.kafkaProd.SendMessage(saramaMsg)
	if err != nil {
		log.Printf(err.Error())
		return false
	}
	return true
}

func (prod *producer) Close() error {
	return prod.kafkaProd.Close()
}

func prepareMessage(topic, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}
	return msg
}
