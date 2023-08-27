package xkafka

import (
	"context"
	"log"
	"testing"

	"github.com/IBM/sarama"
)

func TestNewConsumerHubPlain(t *testing.T) {
	hub, err := NewConsumerHub(&Config{
		Url:        "localhost:9092",
		SASLEnable: false,
		User:       "",
		Password:   "",
	})
	if err != nil {
		panic(err)
	}
	hub.Add(exampleConsumerGroupHandler{})
	go func() {
		hub.Start(context.Background())
	}()
	<-make(chan struct{})
}

func TestNewConsumerHubScram(t *testing.T) {
	hub, err := NewConsumerHub(&Config{
		Url:        "localhost:9092",
		SASLEnable: true,
		User:       "",
		Password:   "",
		Mechanism:  "SCRAM-SHA-256",
	})
	if err != nil {
		panic(err)
	}
	hub.Add(exampleConsumerGroupHandler{})
	go func() {
		hub.Start(context.Background())
	}()
	<-make(chan struct{})
}

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Topic() string {
	return "local.go.test"
}
func (exampleConsumerGroupHandler) GroupId() string {
	return "local.go.testgroupid"
}
func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (exampleConsumerGroupHandler) ConsumeClaim(cgs sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := cgs.Context()
	for msg := range claim.Messages() {
		process(ctx, msg)
		cgs.MarkMessage(msg, "")
	}
	return nil
}

func process(ctx context.Context, msg *sarama.ConsumerMessage) error {
	log.Printf("message topic: %s, partition: %d, offset: %d, key: %s, value: %s", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
	return nil
}
