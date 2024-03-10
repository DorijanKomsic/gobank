package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type ApiServer struct {
    listenAddr string
    store Storage
}

func NewApiServer(listenAddr string, store Storage) *ApiServer {
    return &ApiServer{
        listenAddr: listenAddr,
        store: store,
    }
}

func (s *ApiServer) Run() {
    router := mux.NewRouter()

    router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
    router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccountByID))

    log.Println("JSON API server running on port: ", s.listenAddr)
    http.ListenAndServe(s.listenAddr, router)
}

func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
    switch r.Method {
    case "GET":
        return s.handleGetAccount(w, r)
    case "POST":
        return s.handleCreateAccount(w, r)
    case "DELETE":
        return s.handleDeleteAccount(w, r)
    }
    return fmt.Errorf("Method not allowed %s", r.Method)
} 

func (s *ApiServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
    id := mux.Vars(r)["id"]
    fmt.Println(id)
    return WriteJson(w, http.StatusOK, &Account{})
} 

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
    accounts, err := s.store.GetAccounts() 
    if err != nil {
        return err
    }
    return WriteJson(w, http.StatusOK, accounts)
}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
    createAccountReq := new(CreateAccountRequest)
    if err := json.NewDecoder(r.Body).Decode(createAccountReq); err != nil {
       return err
    }

    account := NewAccount(createAccountReq.FirstName, createAccountReq.LastName)

    if err := s.store.CreateAccount(account); err != nil {
        return err
    }

    return WriteJson(w, http.StatusOK, account)
} 

func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
    return nil
} 

func (s *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
    return nil
} 

func WriteJson(w http.ResponseWriter, status int, v any) error {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(status)
    return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error
type ApiError struct {
    Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if err := f(w, r); err != nil {
            WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
        }
    }
}
