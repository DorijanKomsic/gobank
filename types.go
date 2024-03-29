package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
}

type Account struct {
    ID int `json:"id"`
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    Number int64 `json:"number"`
    Balance float64 `json:"balance"`
    CreatedAt time.Time `json:"createdAt"`
}

type TransferRequest struct {
    ToAccount int `json:"toAccount"`
    Amount float64 `json:"amount"`
}

func NewAccount(firstName, lastName string) *Account {
    return &Account{
        FirstName: firstName,
        LastName: lastName,
        Number: int64(rand.Intn(1000000)),
        CreatedAt: time.Now().UTC(),
    }
}
