package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
    CreateAccount(*Account) error
    DeleteAccount(int) error
    UpdateAccount(*Account) error
    GetAccounts() ([]*Account, error)
    GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
    db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
    connStr := "user=postgres password=debug dbname=postgres sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &PostgresStore{
        db: db,
    }, nil }

func (s *PostgresStore) Init() error {
    return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
    query := `CREATE TABLE IF NOT EXISTS account (
        id serial primary key,
        first_name varchar(50),
        last_name varchar(50),
        number serial,
        balance float,
        created_at timestamp
    )`

    _, err := s.db.Exec(query)
    return err
}

func (s *PostgresStore) CreateAccount(account *Account) error {
    execQuery := `
    INSERT INTO account (first_name, last_name, number, balance, created_at)
    VALUES ($1, $2, $3, $4, $5)
    `
    resp, err := s.db.Exec(
    execQuery, 
    account.FirstName, 
    account.LastName, 
    account.Number,
    account.Balance,
    account.CreatedAt)
    
    if err != nil {
        return err
    }

    fmt.Printf("%+v\n", resp)

    return nil 
} 

func (s *PostgresStore) DeleteAccount(id int) error {
    return nil
}
func (s *PostgresStore) UpdateAccount(*Account) error {
    return nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
    rows, err := s.db.Query("SELECT * FROM account")
    if err != nil {
        return nil, err
    }
    accounts := []*Account{}
    for rows.Next() {
        account := new(Account)
        err := rows.Scan(
            &account.ID, 
            &account.FirstName, 
            &account.LastName, 
            &account.Number, 
            &account.Balance, 
            &account.CreatedAt)
            
        if err != nil {
            return nil, err
        }
        accounts = append(accounts, account)
    }
    return accounts, nil
} 

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
    return nil, nil
}
