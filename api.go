package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
    w.WriteHeader(status)
    w.Header().Set("Content-Type", "application/json")
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

type ApiServer struct {
    listenAddr string
}

func NewApiServer(listenAddr string) *ApiServer {
    return &ApiServer{
        listenAddr: listenAddr,
    }
}

func (s *ApiServer) Run() {
    router := mux.NewRouter()

    router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
    router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))

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
func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
    id := mux.Vars(r)["id"]
    fmt.Println(id)
    return WriteJson(w, http.StatusOK, &Account{})
} 
func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
    return nil
} 
func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
    return nil
} 
func (s *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
    return nil
} 
