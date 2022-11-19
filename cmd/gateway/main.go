package main

import (
	"context"
	"github.com/h-celel/friendly-spoon/internal/api/healthcheck"
	"github.com/h-celel/friendly-spoon/internal/config"
	"github.com/h-celel/friendly-spoon/internal/reverseproxy"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	catchShutdown(cancel)
	env := config.NewEnvironment()
	log.Println("starting gateway...")

	reverseproxy.Init(ctx, cancel, env)
	healthcheck.Init(ctx, cancel, env)

	<-ctx.Done()
}

func catchShutdown(cancel context.CancelFunc) {
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-cancelChan
		log.Printf("Caught SIGTERM %v", sig)
		cancel()
	}()
}
