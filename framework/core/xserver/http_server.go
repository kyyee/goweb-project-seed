package xserver

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type (
	Config struct {
		Port int    `json:"port" mapstructure:"port"`
		Mode string `json:"mode" mapstructure:"mode"`
	}
	Http struct {
		ge  *gin.Engine
		c   *Config
		nhs *http.Server
	}
	option struct {
		port int
	}
	Option func(o *option)
)

func New(c *Config) *Http {
	switch c.Mode {
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	case gin.TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
	e := gin.New()
	if gin.DebugMode == c.Mode || gin.TestMode == c.Mode {
		e.Use(Debug(c))
	}
	// router.GET("/introspect/prometheus/stat", )
	e.GET("/healthcheck", func(c *gin.Context) {
		c.SecureJSON(http.StatusOK, nil)
	})

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", c.Port),
		Handler: e,
	}
	return &Http{
		ge:  e,
		c:   c,
		nhs: s,
	}
}

func (h *Http) Start(ctx context.Context) error {
	log.Printf("[start] http server start, listen on %d", h.c.Port)
	return h.nhs.ListenAndServe()
}

func (h *Http) Stop(ctx context.Context) error {
	log.Printf("[stop] http server stop")
	return h.nhs.Shutdown(ctx)
}

func (h *Http) Router() *gin.Engine {
	return h.ge
}

func WithPort(port int) Option {
	return func(o *option) { o.port = port }
}
