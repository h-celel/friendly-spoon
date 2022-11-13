package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func (h *handler) callbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("state") != h.sessions.GetString(r, "state") {
		_, _ = fmt.Fprint(w, "Invalid state parameter.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accessTypeCodeVerifier := oauth2.SetAuthURLParam("code_verifier", h.sessions.GetString(r, "verifier"))

	token, err := h.config.Exchange(
		r.Context(),
		r.URL.Query().Get("code"),
		oauth2.AccessTypeOffline,
		accessTypeCodeVerifier,
	)
	if err != nil {
		log.Println(err)
		_, _ = fmt.Fprint(w, "Failed to convert an authorization code into a token.")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	idToken, err := h.verifyIDToken(r.Context(), token)
	if err != nil {
		log.Println(err)
		_, _ = fmt.Fprint(w, "Failed to verify ID Token.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var profile map[string]interface{}
	if err = idToken.Claims(&profile); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if rawProfile, err := json.Marshal(profile); err == nil {
		h.sessions.Put(r, "profile", string(rawProfile))
	} else {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	h.sessions.Put(r, "access_token", token.AccessToken)
	h.sessions.Put(r, "refresh_token", token.RefreshToken)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (h *handler) verifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: h.config.ClientID,
	}

	return h.provider.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}
