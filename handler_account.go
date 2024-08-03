package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
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

func (apiCfg *apiConfig) handlerGetAccount(w http.ResponseWriter, r *http.Request, account database.Account) {
	respondWithJSON(w, 200, dbAccountToAccount(account))
}

func (apiCfg *apiConfig) handlerGetPostsForAccount(w http.ResponseWriter, r *http.Request, account database.Account) {
	posts, err := apiCfg.DB.GetPostsForAccount(r.Context(), database.GetPostsForAccountParams{
		AccountID: account.ID,
		Limit:     10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't get posts: %v", err))
	}

	respondWithJSON(w, 200, dbPostsToPosts(posts))
}
