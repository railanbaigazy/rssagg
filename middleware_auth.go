package main

import (
	"fmt"
	"net/http"

	"github.com/railanbaigazy/rssagg/internal/auth"
	"github.com/railanbaigazy/rssagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.Account)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("auth error: %v", err))
			return
		}
		account, err := apiCfg.DB.GetAccountByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("couldn't get account: %v", err))
			return
		}
		handler(w, r, account)
	}
}
