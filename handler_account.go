package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/railanbaigazy/rssagg/internal/auth"
	"github.com/railanbaigazy/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateAccount(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}

	account, err := apiCfg.DB.CreateAccount(r.Context(), database.CreateAccountParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't create account: %v", err))
		return
	}

	respondWithJSON(w, 201, dbAccountToAccount(account))
}

func (apiCfg *apiConfig) handlerGetAccount(w http.ResponseWriter, r *http.Request) {
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

	respondWithJSON(w, 200, dbAccountToAccount(account))
}
