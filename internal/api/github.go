package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/8Whoknow3/github-activity-cli/internal/models"
)

const baseURL = "https://api.github.com"

func (c *Client) GetUserEvents(username string) ([]models.Event, error) {
	api := fmt.Sprintf("%s/users/%s/events", baseURL, username)

	req, err := http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "github-activity-cli")
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("github user %q not found", username)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned unexpected status %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)

	var events []models.Event

	if err := decoder.Decode(&events); err != nil {
		return nil, err
	}
	return events, nil
}
