package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/railanbaigazy/rssagg/internal/database"
)

type Account struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func dbAccountToAccount(dbAccount database.Account) Account {
	return Account{
		ID:        dbAccount.ID,
		CreatedAt: dbAccount.CreatedAt,
		UpdatedAt: dbAccount.UpdatedAt,
		Name:      dbAccount.Name,
		APIKey:    dbAccount.ApiKey,
	}
}
