package main

import (
	"context"
	"github.com/golangcollege/sessions"
	"github.com/h-celel/friendly-spoon/internal/api/healthcheck"
	"github.com/h-celel/friendly-spoon/internal/api/rest"
	"github.com/h-celel/friendly-spoon/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	catchShutdown(cancel)
	env := config.NewEnvironment()
	log.Println("starting app...")

	go func() {
		session := sessions.New([]byte(config.SessionsSecret))
		session.Lifetime = config.SessionsLifetime

		rest.Init(ctx, cancel, env, session)
		healthcheck.Init(ctx, cancel, env)
	}()

	<-ctx.Done()
	log.Println("closing app")
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
