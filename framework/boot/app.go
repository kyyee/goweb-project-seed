package boot

import (
	"context"
	"goweb-project-seed/framework/core/xconfig"
	"goweb-project-seed/framework/core/xelasticsearch"
	"goweb-project-seed/framework/core/xgorm"
	"goweb-project-seed/framework/core/xkafka"
	"goweb-project-seed/framework/core/xredis"
	"goweb-project-seed/framework/core/xserver"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	Bootstrap func() *xconfig.Bootstrap

	Config struct {
		Config       any
		Bootstrap    Bootstrap
		Constructor  interface{}
		AssignPath   string
		BannerEnable bool
	}

	App struct {
		bootstrap     *xconfig.Bootstrap
		servers       []xserver.Server
		elasticsearch *xelasticsearch.Client
		producer      *xkafka.Producer
		consumer      *xkafka.ConsumerHub
		redis         *xredis.Client
		db            *gorm.DB
		ge            *gin.Engine
	}

	Middleware gin.HandlerFunc

	RouteRegister     func() error
	ComponentRegister func() error
	CustomizeFunc     func() error
)

func New(c *Config) *App {
	if c.Constructor == nil {
		log.Panicf("please init constructor in App Config")
	}
	if c.Bootstrap == nil {
		log.Panicf("please init boostrap in App Config")
	}
	err := xconfig.Load(c.AssignPath, c.Config)
	if err != nil {
		log.Panicf("load app config failed %v", err)
	}

	return &App{bootstrap: c.Bootstrap()}
}

func (a *App) Route(rr RouteRegister) *App {
	if rr == nil {
		log.Panicf("route register func should not be nil")
	}
	err := rr()
	if err != nil {
		log.Panicf("route register func exec failed, %v", err)
	}
	return a
}

func (a *App) Http(middleware ...gin.HandlerFunc) *App {
	if a.bootstrap.HttpServer == nil {
		log.Panicf("please set http config")
	}
	http := xserver.New(a.bootstrap.HttpServer)
	a.servers = append(a.servers, http)
	router := http.Router()
	router.Use(middleware...)

	a.ge = router
	return a
}

func (a *App) Orm() *App {
	if a.bootstrap == nil && a.bootstrap.Orm == nil {
		log.Panicf("please set orm config")
	}
	db, err := xgorm.New(a.bootstrap.Orm)
	if err != nil {
		log.Panicf("orm init failed")
	}
	a.db = db
	return a
}

func (a *App) Redis() *App {
	if a.bootstrap == nil && a.bootstrap.Redis == nil {
		log.Panicf("please set redis config")
	}
	redis, err := xredis.New(a.bootstrap.Redis)
	if err != nil {
		log.Panicf("redis init failed")
	}
	a.redis = redis
	return a
}

func (a *App) ElasticSearch() *App {
	if a.bootstrap == nil && a.bootstrap.ElasticSearch == nil {
		log.Panicf("please set elastic search config")
	}
	elasticsearch, err := xelasticsearch.New(a.bootstrap.ElasticSearch)
	if err != nil {
		log.Panicf("elastic search init failed")
	}
	a.elasticsearch = elasticsearch
	return a
}

func (a *App) Produce() *App {
	if a.bootstrap == nil && a.bootstrap.Kafka == nil {
		log.Panicf("please set kafka config")
	}
	producer, err := xkafka.NewProducer(a.bootstrap.Kafka)
	if err != nil {
		log.Panicf("kafka producer init failed")
	}
	a.producer = producer
	return a
}

func (a *App) Consume() *App {
	if a.bootstrap == nil && a.bootstrap.Kafka == nil {
		log.Panicf("please set kafka config")
	}
	consumer, err := xkafka.NewConsumerHub(a.bootstrap.Kafka)
	a.consumer = consumer
	if err != nil {
		log.Panicf("kafka consumer init failed")
	}
	a.servers = append(a.servers, consumer)
	return a
}

func (a *App) Component(f ComponentRegister) *App {
	if f == nil {
		log.Panicf("Component register func should not be nil")
	}
	err := f()
	if err != nil {
		log.Panicf("component exec failed, %v", err)
	}
	return a
}

func (a *App) ServerExtend(s ...xserver.Server) *App {
	a.servers = append(a.servers, s...)
	return a
}

func (a *App) Customize(f CustomizeFunc) *App {
	if f == nil {
		log.Panicf("func should not be nil")
	}
	err := f()
	if err != nil {
		log.Panicf("customize func exec failed, %v", err)
	}
	return a
}

func (a *App) printBanner() {

}

func (a *App) Run() {
	if a.bootstrap.BannerEnable {
		a.printBanner()
	}
	ctx := context.Background()
	for _, s := range a.servers {
		go func(ctx context.Context, s xserver.Server) {
			err := s.Start(ctx)
			if err != nil && err != http.ErrServerClosed {
				log.Panicf("http server start failed %v", err)
			}
		}(ctx, s)
	}
	log.Printf("[boot] all server start ok!")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	<-channel
	for _, s := range a.servers {
		err := s.Stop(ctx)
		if err != nil {
			log.Panicf("http server stop failed %v", err)
		}
	}
}
