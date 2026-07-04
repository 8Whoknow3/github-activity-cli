package main

import (
	"fmt"
	"log"

	client "github.com/8Whoknow3/github-activity-cli/internal/api"
	"github.com/8Whoknow3/github-activity-cli/internal/cli"
	"github.com/8Whoknow3/github-activity-cli/internal/formatter"
)

func main() {
	args, err := cli.Parse()
	if err != nil {
		log.Fatal(err)
	}

	if args == nil {
		return
	}

	apiClient := client.NewClient()

	events, err := apiClient.GetUserEvents(args.Username)
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range events {
		fmt.Println(formatter.FormatEvent(event))
	}
}
