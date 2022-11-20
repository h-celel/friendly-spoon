package reverseproxy

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/h-celel/friendly-spoon/internal/config"
	"github.com/h-celel/sessions"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Init(_ context.Context, cancel context.CancelFunc, env config.Environment, sessions *sessions.Session) {
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
	h2handler := h2c.NewHandler(sessions.Enable(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("Authorization")) == 0 {
			if accessToken := sessions.GetString(r, "access_token"); len(accessToken) > 0 {
				r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
			}
		}
		proxy.ServeHTTP(w, r)
	})), h2s)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", env.GatewayHostPort),
		Handler: h2handler,
	}

	go func() {
		defer cancel()
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()
}
