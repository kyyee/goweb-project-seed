package utils

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnowFlakeId(t *testing.T) {
	sf, err := NewSnowFlake(12)
	assert.Nil(t, err)
	id := sf.Generate()
	log.Printf("generate id is: %d", id)
	assert.NotNil(t, id)
}
