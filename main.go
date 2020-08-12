package main

import (
	"context"
	"fmt"

	"github.com/chavacava/lab-mod/myutils"

	"github.com/google/go-github/v32/github"
)

// Fetch all the public organizations' membership of a user.
//
func fetchOrganizations(username string) ([]*github.Organization, error) {
	client := github.NewClient(nil)
	orgs, _, err := client.Organizations.List(context.Background(), username, nil)
	return orgs, err
}

func main() {

	_ = myutils.Do()
	var username string
	fmt.Print("Enter GitHub username: ")
	fmt.Scanf("%s", &username)

	organizations, err := fetchOrganizations(username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for i, organization := range organizations {
		fmt.Printf("%v. %v\n", i+1, organization.GetLogin())
	}
}
