package rest

import (
	"context"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/mux"
	"github.com/h-celel/friendly-spoon/internal/config"
	"github.com/h-celel/sessions"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func Init(_ context.Context, cancel context.CancelFunc, env config.Environment, sessions *sessions.Session) {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+env.Auth0Domain+"/",
	)
	if err != nil {
		log.Println(err)
		cancel()
		return
	}

	conf := oauth2.Config{
		ClientID:     env.Auth0ClientID,
		ClientSecret: env.Auth0ClientSecret,
		RedirectURL:  env.Auth0CallbackURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}
	conf.Endpoint.AuthStyle = oauth2.AuthStyleInHeader

	handler := newHandler(env, provider, conf, sessions)

	router := mux.NewRouter()

	router.Path("/login").Methods(http.MethodGet).HandlerFunc(handler.loginHandler)
	router.Path("/callback").Methods(http.MethodGet).HandlerFunc(handler.callbackHandler)
	router.Path("/logout").Methods(http.MethodGet).HandlerFunc(handler.logoutHandler)

	go func() {
		defer cancel()
		err := http.ListenAndServe(fmt.Sprintf(":%d", env.RESTPort), sessions.Enable(router))
		if err != nil {
			log.Println(err)
		}
	}()
}

type handler struct {
	env      config.Environment
	provider *oidc.Provider
	config   oauth2.Config
	sessions *sessions.Session
}

func newHandler(env config.Environment, provider *oidc.Provider, config oauth2.Config, sessions *sessions.Session) *handler {
	return &handler{env: env, provider: provider, config: config, sessions: sessions}
}
