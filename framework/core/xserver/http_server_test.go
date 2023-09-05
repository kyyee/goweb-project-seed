package xserver

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestNewHttpServer(t *testing.T) {
	server := New(&Config{Port: 9093, Mode: "debug"})
	ctx := context.Background()
	go func() {
		err := server.Start(ctx)
		if err != nil && err != http.ErrServerClosed {
			t.Error(err)
		}
	}()
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	<-channel
	err := server.Stop(ctx)
	if err != nil {
		t.Error(err)
	}
}
