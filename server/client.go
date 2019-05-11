package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ChimeraCoder/anaconda"
)

type Client struct {
	// http.Client
	TwitterAPI *anaconda.TwitterApi
	UserData   map[string]string
}

func (*Client) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		defer r.Body.Close()
		log.Println("hello from golang in post method")
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&client.UserData)
		if err != nil {
			log.Fatal("decoder error")
			panic(err)
		}
		w.Header().Set("Server", "A Go Web Server")
		w.WriteHeader(200)
	case "GET":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		defer r.Body.Close()
		log.Println("hello from golang in get method")
		encoder := json.NewEncoder(w)
		encoder.Encode(getVerifiedUserObjects(client.UserData["query"]))
		w.Header().Set("Server", "A Go Web Server")
		w.WriteHeader(200)
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}

func getVerifiedUserObjects(target string) *[]anaconda.User {
	var verifiedUsers []anaconda.User
	users, err := client.TwitterAPI.GetUserSearch(target, nil)
	if err != nil {
		log.Fatal("twitter eror")
		panic(err)
	}

	for _, user := range users {
		if user.Verified == true {
			verifiedUsers = append(verifiedUsers, user)
		}
	}

	return &verifiedUsers
}
