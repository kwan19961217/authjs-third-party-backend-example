package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	EmailVerified time.Time `json:"emailVerified"`
	Image         string    `json:"image"`
}

type Account struct {
	Id                string `json:"id"`
	UserId            string `json:"userId"`
	Provider          string `json:"provider"`
	ProviderAccountId string `json:"providerAccountId"`
	RefreshToken      string `json:"refresh_token"`
	AccessToken       string `json:"access_token"`
	ExpiresAt         int    `json:"expires_at"`
	TokenType         string `json:"token_type"`
	Scope             string `json:"scope"`
	IdToken           string `json:"id_token"`
	SessionState      string `json:"session_state"`
}

func main() {
	var users []User
	var accounts []Account

	app := http.NewServeMux()
	app.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(user)
		user.Id = strconv.Itoa(len(users) + 1)
		users = append(users, user)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	})
	app.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		var user User
		for _, u := range users {
			if u.Id == id {
				user = u
				break
			}
		}
		if user.Id == "" {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
	})
	app.HandleFunc("GET /users/email/{email}", func(w http.ResponseWriter, r *http.Request) {
		email := r.PathValue("email")
		var user User
		for _, u := range users {
			if u.Email == email {
				user = u
				break
			}
		}
		if user.Id == "" {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
	})
	app.HandleFunc("GET /users/account/{provider}/{providerAccountId}", func(w http.ResponseWriter, r *http.Request) {
		provider := r.PathValue("provider")
		providerAccountId := r.PathValue("providerAccountId")
		var account Account
		for _, a := range accounts {
			if a.Provider == provider && a.ProviderAccountId == providerAccountId {
				account = a
				break
			}
		}
		if account.Id == "" {
			http.Error(w, "account not found", http.StatusNotFound)
			return
		}
		var user User
		for _, u := range users {
			if u.Id == account.UserId {
				user = u
				break
			}
		}
		if user.Id == "" {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
	})

	app.HandleFunc("PUT /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var found bool
		for i, u := range users {
			if u.Id == id {
				users[i] = user
				found = true
				break
			}
		}
		if !found {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
	})

	app.HandleFunc("POST /accounts", func(w http.ResponseWriter, r *http.Request) {
		var account Account
		if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		account.Id = strconv.Itoa(len(accounts) + 1)
		accounts = append(accounts, account)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(account)
	})

	log.Fatal(http.ListenAndServe(":8080", app))
}
