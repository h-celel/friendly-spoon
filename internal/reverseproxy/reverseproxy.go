package reverseproxy

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/h-celel/friendly-spoon/internal/config"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Init(_ context.Context, cancel context.CancelFunc, env config.Environment) {
	targetUrl, err := url.Parse(env.GatewayTargetURL)
	if err != nil {
		fmt.Println(err)
		cancel()
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetUrl)
	proxy.Transport = &http2.Transport{
		AllowHTTP: true,
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			return net.Dial(network, addr)
		},
	}

	h2s := &http2.Server{}
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", env.GatewayHostPort),
		Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			proxy.ServeHTTP(w, r)
		}), h2s),
	}

	go func() {
		defer cancel()
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()
}
