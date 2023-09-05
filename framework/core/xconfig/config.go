package xconfig

import (
	"goweb-project-seed/framework/core/xdatasource"
	"goweb-project-seed/framework/core/xelasticsearch"
	"goweb-project-seed/framework/core/xerrors"
	"goweb-project-seed/framework/core/xkafka"
	"goweb-project-seed/framework/core/xredis"
	"log"
	"reflect"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Bootstrap struct {
	DataSource    *xdatasource.Config    `json:"data_source" mapstructure:"data_source"`
	Kafka         *xkafka.Config         `json:"kafka" mapstructure:"kafka"`
	Redis         *xredis.Config         `json:"redis" mapstructure:"redis"`
	ElasticSearch *xelasticsearch.Config `json:"elastic_search" mapstructure:"elastic_search"`
}

func Load(path string, value any) error {
	if len(path) <= 0 {
		log.Fatalf("config path can not empty")
		return xerrors.NewAssignCodeError(xerrors.GET_CONFIGURATION_ERROR)
	}
	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		log.Fatalf("get config failed, config should be pointer")
		return xerrors.NewAssignCodeError(xerrors.GET_CONFIGURATION_ERROR)
	}

	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return errors.WithStack(err)
	}
	err = viper.Unmarshal(value)

	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
