package formatter

import (
	"fmt"

	"github.com/8Whoknow3/github-activity-cli/internal/models"
)

// FormatEvent formats a GitHub event into a human-readable message.
func FormatEvent(event models.Event) string {
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
		if event.Payload.Action != "" {
			return fmt.Sprintf(
				"%s issue in %s",
				event.Payload.Action,
				event.Repo.Name,
			)
		}

		return fmt.Sprintf(
			"Issue activity in %s",
			event.Repo.Name,
		)

	case "CreateEvent":
		if event.Payload.Ref != "" {
			return fmt.Sprintf(
				"Created %s %q in %s",
				event.Payload.RefType,
				event.Payload.Ref,
				event.Repo.Name,
			)
		}

		return fmt.Sprintf(
			"Created %s in %s",
			event.Payload.RefType,
			event.Repo.Name,
		)

	default:
		return fmt.Sprintf(
			"%s in %s",
			event.Type,
			event.Repo.Name,
		)
	}
}
