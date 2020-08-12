package utils

import "github.com/google/go-github/v22/github"

func Do() *github.Client {
	return github.NewClient(nil)
}