package cli

import (
	"fmt"
	"os"
)

type Args struct {
	Username string
}

func Parse() (*Args, error) {
	if len(os.Args) < 2 {
		printHelp()
		return nil, fmt.Errorf("missing GitHub username")
	}

	switch os.Args[1] {
	case "help", "-h", "--help":
		printHelp()
		return nil, nil
	}

	return &Args{
		Username: os.Args[1],
	}, nil
}

func printHelp() {
	fmt.Println(`GitHub Activity CLI

Fetch and display the recent public activity of a GitHub user.

Usage:
    github-activity <username>

Arguments:
    username    GitHub username

Options:
    -h, --help  Show this help message

Examples:
    github-activity octocat
    github-activity torvalds
`)
}
