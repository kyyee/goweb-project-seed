package xkafka

import (
	"context"
	"errors"
	"goweb-project-seed/framework/core/xgoroutine"
	"log"
	"strings"

	"github.com/IBM/sarama"
)

type Consumer interface {
	Topic() string
	GroupId() string
	sarama.ConsumerGroupHandler
}

type ConsumerHub struct {
	consumers []Consumer
	cfg       *sarama.Config
	brokers   []string
	groups    []sarama.ConsumerGroup
}

func NewConsumerHub(c *Config) (*ConsumerHub, error) {
	if len(c.Url) <= 0 {
		log.Fatalf("kafka url should not be empty")
		return nil, errors.New("kafka url should not be empty")
	}
	cfg := sarama.NewConfig()
	cfg.Net.SASL.Enable = c.SASLEnable
	cfg.Net.SASL.User = c.User
	cfg.Net.SASL.Password = c.Password // todo 密码使用enc加密
	cfg.Consumer.Return.Errors = true
	cfg.Consumer.Offsets.Initial = sarama.OffsetNewest
	cfg.Version = sarama.V3_3_1_0
	c.Mechanism = strings.ToUpper(c.Mechanism)
	if c.Mechanism == sarama.SASLTypeSCRAMSHA512 {
		cfg.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
		cfg.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
			return &XDGSCRAMClient{HashGeneratorFcn: SHA512}
		}
	} else if c.Mechanism == sarama.SASLTypeSCRAMSHA256 {
		cfg.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256
		cfg.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
			return &XDGSCRAMClient{HashGeneratorFcn: SHA256}
		}
	} else {
		cfg.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	}
	return &ConsumerHub{
		cfg:     cfg,
		brokers: strings.Split(c.Url, ","),
	}, nil
}

func (hub *ConsumerHub) Add(c Consumer) {
	hub.consumers = append(hub.consumers, c)
}

func (hub *ConsumerHub) Start(ctx context.Context) error {
	log.Printf("[start] kafka consumers start")
	for _, c := range hub.consumers {
		group, err := sarama.NewConsumerGroup(hub.brokers, c.GroupId(), hub.cfg)
		if err != nil {
			return err
		}
		hub.groups = append(hub.groups, group)
		go func(g sarama.ConsumerGroup, handler Consumer) {
			defer xgoroutine.Recover()
			for {
				if err := g.Consume(ctx, []string{handler.Topic()}, handler); err != nil {
					log.Fatalf("xkafka consume failed, topic: %s, group: %s, err: %v", handler.Topic(), handler.GroupId(), err)
				}
				if ctx.Err() != nil {
					log.Fatalf("xkafka consume ctx error, topic: %s, group: %s, err: %v", handler.Topic(), handler.GroupId(), err)
					break
				}
			}
		}(group, c)
	}
	return nil
}

func (hub *ConsumerHub) Stop(ctx context.Context) error {
	for _, group := range hub.groups {
		group.Close()
	}

	log.Printf("[start] kafka consumers stop")
	return nil
}
