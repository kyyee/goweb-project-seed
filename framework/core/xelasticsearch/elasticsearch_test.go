package xelasticsearch

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client, err := New(&Config{
		Addresses: []string{"http://localhost:9200"},
		Username:  "elastic",
		Password:  "123456",
	})
	assert.Nil(t, err)
	assert.NotNil(t, client)
	result, err := client.Client().Search(
		client.Client().Search.WithContext(context.Background()),
		client.Client().Search.WithIndex("monitoring-es-8-2023.08.26"),
		client.Client().Search.WithTrackTotalHits(true),
		client.Client().Search.WithPretty(),
	)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	log.Printf("elasticsearch result: %s", result)
}
