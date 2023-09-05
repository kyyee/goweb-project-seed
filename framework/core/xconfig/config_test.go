package xconfig

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	var c AppConfig
	err := Load("config.yaml", &c)
	assert.Nil(t, err)
	log.Printf("config: %+v", c)
	assert.Equal(t, "gps", c.Name)
	assert.Equal(t, "gps", c.Bootstrap.DataSource.DbName)
	assert.Equal(t, "plain", c.Bootstrap.Kafka.Mechanism)

}

type AppConfig struct {
	Name      string    `mapstrucure:"name"`
	Bootstrap Bootstrap `mapstrucure:"bootstrap"`
}
