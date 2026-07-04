package formatter

import (
	"fmt"

	"github.com/8Whoknow3/github-activity-cli/internal/models"
)

func Format(event models.Event) string {

	switch event.Type {

	case "PushEvent":
		return fmt.Sprintf(
			"Pushed %d commits to %s",
			len(event.Payload.Commits),
			event.Repo.Name,
		)

	case "WatchEvent":
		return fmt.Sprintf(
			"Starred %s",
			event.Repo.Name,
		)

	case "ForkEvent":
		return fmt.Sprintf(
			"Forked %s",
			event.Repo.Name,
		)

	case "IssuesEvent":
		return fmt.Sprintf(
			"Opened an issue in %s",
			event.Repo.Name,
		)

	case "CreateEvent":
		return fmt.Sprintf(
			"Created %s %q in %s",
			event.Payload.RefType,
			event.Payload.Ref,
			event.Repo.Name,
		)

	default:
		return fmt.Sprintf(
			"%s on %s",
			event.Type,
			event.Repo.Name,
		)
	}
}
