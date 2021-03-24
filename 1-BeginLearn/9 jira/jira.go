package main

import (
	"fmt"
	jira "github.com/andygrunwald/go-jira"
)
func main() {
	tp := jira.BasicAuthTransport{
		Username: "grp.cmqa1",
		Password: "5tg@Beijing01",
	}

	client, err := jira.NewClient(tp.Client(), "http://jira.ws.netease.com")

	u, _, err := client.User.Get("some_user")
	fmt.Println(err)

	fmt.Printf("\nEmail: %v\nSuccess!\n", u.EmailAddress)
}