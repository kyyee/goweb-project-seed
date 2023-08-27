package xkafka

import (
	"goweb-project-seed/framework/core/utils"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProducerHubPlain(t *testing.T) {
	_, err := NewProducer(&Config{
		Url:        "localhost:9092",
		SASLEnable: false,
		User:       "",
		Password:   "",
		Mechanism:  "plain",
	})
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
}

func TestNewProducerHubScram(t *testing.T) {
	_, err := NewProducer(&Config{
		Url:        "localhost:9092",
		SASLEnable: true,
		User:       "",
		Password:   "",
		Mechanism:  "SCRAM-SHA-256",
	})
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
}

func TestSendMsg(t *testing.T) {
	p, err := NewProducer(&Config{
		Url:        "localhost:9092",
		SASLEnable: false,
		User:       "",
		Password:   "",
		Mechanism:  "SCRAM-SHA-256",
	})
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
	type Res struct {
		Code string      `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
	res := &Res{
		Code: "0000000000",
		Msg:  "操作成功",
		Data: "test",
	}
	partiton, offset, err := p.Send("local.go.test", "", utils.ToJsonStr(res))
	if err != nil {
		panic(err)
	}
	log.Printf("produce message: %s, partition: %d, offset: %d", utils.ToJsonStr(res), partiton, offset)
}
