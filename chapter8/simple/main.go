package main

import (
	"encoding/json"
	"golang.org/x/net/context"
	"log"
	"net/http"

	"github.com/google/go-github/v38/github"
)

func main() {
	client := github.NewClient(&http.Client{})

	ctx := context.Background()
	repo, _, err := client.Repositories.Get(ctx, "golang", "go")

	if err != nil {
		log.Fatalf(err.Error())
	}

	r, err := json.MarshalIndent(repo, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Println(string(r))

}
