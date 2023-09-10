package boot

import (
	"goweb-project-seed/framework/core/xconfig"
	"goweb-project-seed/framework/core/xelasticsearch"
	"goweb-project-seed/framework/core/xkafka"
	"goweb-project-seed/framework/core/xredis"
	"log"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestAppRun(t *testing.T) {
	ct := &ConfigTemplate{}
	app := New(&Config{
		Bootstrap: func() *xconfig.Bootstrap {
			return ct.Bootstrap
		},
		Constructor: func() *ConfigTemplate {
			return ct
		},
		Config:     ct,
		AssignPath: "config.yaml",
	})
	app.Http().Orm().Redis().Produce().Consume().ElasticSearch().Component(Component).Route(Route).Run()
}

func Route() error {
	RouteRegisterTest
}

func RouteRegisterTest(ge *gin.Engine, api *ApiTest) {
	ge.GET("/test", func(ctx *gin.Context) {
		api.HelloWorld(ctx)
	})
}

type RepoTest struct {
	db            *gorm.DB
	redis         *xredis.Client
	elasticsearch *xelasticsearch.Client
	producer      *xkafka.Producer
	Consumer      *xkafka.Consumer
}

type ApiTest struct {
	repo *RepoTest
}

func NewApiTest(repo *RepoTest) *ApiTest {
	return &ApiTest{
		repo: repo,
	}
}

type ConfigTemplate struct {
	Bootstrap *xconfig.Bootstrap `mapstructure:"bootstarp"`
}

func (a *ApiTest) HelloWorld(ctx *gin.Context) {
	log.Printf("hello %s", "world")
	ctx.String(200, "hello world")
}
