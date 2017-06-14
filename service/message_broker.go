package service

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
)

type MessageBroker interface {
	Send(message interface{})
	Close() error
}

func NewKafkaBroker(brokers []string, retries int, topic string) (MessageBroker, error) {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = retries
	config.Producer.RequiredAcks = sarama.WaitForAll
	producer, err := sarama.NewAsyncProducer(brokers, config)
	broker := &kafkaBroker{producer, topic}
	if err != nil {
		return nil, err
	}
	return broker, err
}

type kafkaBroker struct {
	producer sarama.AsyncProducer
	topic    string
}

func (b *kafkaBroker) Send(obj interface{}) {
	message, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
	}
	msg := &sarama.ProducerMessage{
		Topic: b.topic,
		//Key:   sarama.StringEncoder(strTime),
		Value: sarama.StringEncoder(message),
	}
	b.producer.Input() <- msg
}

func (b *kafkaBroker) Close() error {
	return b.producer.Close()
}
