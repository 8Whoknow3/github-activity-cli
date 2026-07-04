package models

type Event struct {
	Type    string  `json:"type"`
	Repo    Repo    `json:"repo"`
	Payload Payload `json:"payload"`
}

type Repo struct {
	Name string `json:"name"`
}

type Payload struct {
	Action      string      `json:"action"`
	RefType     string      `json:"ref_type"`
	Ref         string      `json:"ref"`
	Commits     []Commit    `json:"commits"`
	Issue       Issue       `json:"issue"`
	PullRequest PullRequest `json:"pull_request"`
}

type Commit struct {
	Message string `json:"message"`
}

type Issue struct {
	Number int `json:"number"`
}

type PullRequest struct {
	Number int `json:"number"`
}
