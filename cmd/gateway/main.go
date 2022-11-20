package main

import (
	"context"
	"encoding/base64"
	"github.com/h-celel/friendly-spoon/internal/api/healthcheck"
	"github.com/h-celel/friendly-spoon/internal/api/reverseproxy"
	"github.com/h-celel/friendly-spoon/internal/config"
	"github.com/h-celel/sessions"
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

	key, _ := base64.StdEncoding.DecodeString(config.SessionsSecret)
	session := sessions.New(key)
	session.Lifetime = config.SessionsLifetime

	reverseproxy.Init(ctx, cancel, env, session)
	healthcheck.Init(ctx, cancel, env)

	<-ctx.Done()
	log.Println("closing gateway")
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
