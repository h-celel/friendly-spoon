package healthcheck

import (
	"context"
	"fmt"
	"github.com/h-celel/friendly-spoon/internal/config"
	"log"
	"net/http"
	"time"
)

func Init(_ context.Context, cancel context.CancelFunc, env config.Environment) {
	mux := http.NewServeMux()

	mux.Handle("/status", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintf(w, "%s - time: '%v'", config.AppName, time.Now())
	}))

	go func() {
		defer cancel()
		err := http.ListenAndServe(fmt.Sprintf(":%d", env.HealthcheckPort), mux)
		if err != nil {
			log.Println(err)
		}
	}()
}
