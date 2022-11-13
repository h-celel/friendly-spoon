package rest

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func (h *handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	state, err := generateRandomState()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	verifier, err := generatePKCECodeVerifier()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.sessions.Put(r, "state", state)
	h.sessions.Put(r, "verifier", verifier)

	challenge := generatePKCEChallenge(verifier)

	accessTypeCodeChallenge := oauth2.SetAuthURLParam("code_challenge", challenge)
	accessTypeCodeChallengeMethod := oauth2.SetAuthURLParam("code_challenge_method", "S256")

	authCodeURL := h.config.AuthCodeURL(
		state,
		oauth2.AccessTypeOffline,
		accessTypeCodeChallenge,
		accessTypeCodeChallengeMethod,
	)
	http.Redirect(w, r, authCodeURL, http.StatusTemporaryRedirect)
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

func generatePKCECodeVerifier() (string, error) {
	verifier := make([]byte, 32)
	_, err := rand.Read(verifier)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(verifier[:]), nil
}

func generatePKCEChallenge(verifier string) string {
	challenge := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(challenge[:])
}
