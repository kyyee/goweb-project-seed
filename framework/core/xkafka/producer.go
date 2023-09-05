package xkafka

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/IBM/sarama"
)

type Producer struct {
	BrokerServers string
	Config        *sarama.Config
	SyncProducer  sarama.SyncProducer
	ASyncProducer sarama.AsyncProducer
}

func NewProducer(c *Config) (*Producer, error) {
	if len(c.Url) <= 0 {
		log.Fatalf("kafka url should not be empty")
		return nil, errors.New("kafka url should not be empty")
	}
	cfg := sarama.NewConfig()
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Return.Successes = true
	cfg.Net.SASL.Enable = c.SASLEnable
	cfg.Net.SASL.User = c.User
	cfg.Net.SASL.Password = c.Password // todo 密码使用enc加密
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

	client, err := sarama.NewSyncProducer(strings.Split(c.Url, ","), cfg)
	if err != nil {
		return nil, err
	}
	return &Producer{
		BrokerServers: c.Url,

		Config:       cfg,
		SyncProducer: client,
	}, nil
}

func (p *Producer) Send(topic, key string, data string) (partition int32, offset int64, err error) {
	if p.SyncProducer == nil {
		return 0, 0, errors.New("kafka sync producer should not be nil")
	}
	pm := &sarama.ProducerMessage{}
	pm.Topic = topic
	pm.Key = sarama.StringEncoder(key)
	pm.Value = sarama.StringEncoder(data)
	pm.Timestamp = time.Now()
	return p.SyncProducer.SendMessage(pm)
}
