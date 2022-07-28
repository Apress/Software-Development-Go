package main

import (
	"context"
	"github.com/shurcooL/githubv4"
	rt "graphql/roundtripper"
	"log"
	"net/http"
	"os"
)

type graphqlData struct {
	Repository struct {
		CreatedAt githubv4.DateTime
		ForkCount githubv4.Int
		Labels    struct {
			Edges []struct {
				Node struct {
					Name githubv4.String
				}
			}
		} `graphql:"labels(first: $labelcount)"`
		Issues struct {
			Edges []struct {
				Node struct {
					Title githubv4.String
				}
			}
		} `graphql:"issues(first: $issuescount)"`
		CommitComments struct {
			TotalCount githubv4.Int
			Edges      []struct {
				Node struct {
					Author struct {
						URL   githubv4.String
						Login githubv4.String
					}
				}
			}
		} `graphql:"commitComments(first: $commitcount)"`
	} `graphql:"repository(owner: $owner, name: $name) "`
	RateLimit struct {
		Cost *int
	}
}

func main() {
	t, f := os.LookupEnv("GITHUB_TOKEN")

	if !f {
		log.Fatalf("Cannot find GITHUB_TOKEN env variable")
	}
	graphClient := githubv4.NewClient(&http.Client{
		Transport: rt.NewTransport(t),
	})

	data := new(graphqlData)
	vars := map[string]interface{}{
		"owner":       githubv4.String("golang"),
		"name":        githubv4.String("go"),
		"labelcount":  githubv4.Int(10),
		"issuescount": githubv4.Int(10),
		"commitcount": githubv4.Int(10),
	}
	if err := graphClient.Query(context.Background(), data, vars); err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Total number of fork : ", data.Repository.ForkCount)
	log.Println("Total number of labels : ", len(data.Repository.Labels.Edges))

	log.Println("----------------------------------")
	for _, issue := range data.Repository.Issues.Edges {
		log.Printf("Issue title - %s ", issue.Node.Title)
	}
	log.Println("----------------------------------")
	for _, commit := range data.Repository.CommitComments.Edges {
		log.Printf("Commit author (%s), url (%s) ", commit.Node.Author.Login, commit.Node.Author.URL)
	}
}
