package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/ntk148v/jwt-middleware"
)

const bearerFormat string = "Bearer %s"

type API struct {
	token *jwt.Token
}

func NewAPI() (*API, error) {
	pwd, _ := os.Getwd()
	token, err := jwt.NewToken(jwt.Options{
		PrivateKeyLocation: path.Join(pwd, "keys/test.rsa"),
		PublicKeyLocation:  path.Join(pwd, "keys/test.rsa.pub"),
		SigningMethod:      "RS256",
		TTL:                5 * time.Minute,
		IsBearerToken:      true,
	}, nil)
	if err != nil {
		return nil, err
	}
	return &API{token: token}, nil
}

func (a *API) IssueToken(w http.ResponseWriter, req *http.Request) {
	user, pass, _ := req.BasicAuth()
	if user != "test" || pass != "test" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := make(map[string]interface{})
	data["user"] = user
	tokenString, err := a.token.GenerateToken(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Add("Authorization", fmt.Sprintf(bearerFormat, tokenString))
}

func (A *API) Secret(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Here is the secret")
}

func main() {
	var port = flag.Int("port", 3000, "port to listen on")
	flag.Parse()
	router := mux.NewRouter()
	api, err := NewAPI()
	if err != nil {
		log.Fatal(err)
	}
	pubRouter := router.PathPrefix("/public").Subrouter()
	pubRouter.HandleFunc("/token", api.IssueToken).Methods("POST")
	privRouter := router.PathPrefix("/private").Subrouter()
	privRouter.Use(jwt.Authenticator(api.token))
	privRouter.HandleFunc("/secret", api.Secret).Methods("GET")
	log.Println("Starting the server...")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), router))
}
