# GitHub Activity CLI

A lightweight command-line application written in Go that fetches and displays a GitHub user's recent public activity using the GitHub REST API.

## ✨ Features

- Fetch recent public activity for any GitHub user
- Display activity in a clean, human-readable format
- Supports common GitHub events:
    - PushEvent
    - WatchEvent
    - ForkEvent
    - IssuesEvent
    - CreateEvent
- Graceful error handling
- Lightweight with no third-party dependencies

## 📦 Installation

### Clone the repository

```bash
git clone https://github.com/8Whoknow3/github-activity-cli.git
cd github-activity-cli
```

### Build

```bash
go build -o github-activity ./cmd/github-activity
```

Or run it directly without building:

```bash
go run ./cmd/github-activity <username>
```

## 🚀 Usage

```bash
github-activity <github-username>
```

Example:

```bash
github-activity octocat
```

## Example Output

```text
- Pushed 3 commits to octocat/Hello-World
- Starred octocat/Spoon-Knife
- Forked octocat/git-consortium
- Opened an issue in octocat/Hello-World
- Created branch "feature/login" in octocat/Hello-World
```

## 📁 Project Structure

```
github-activity-cli/
├── cmd/
│   └── github-activity/
│       └── main.go
├── internal/
│   ├── api/
│   ├── cli/
│   ├── formatter/
│   └── models/
├── .gitignore
├── LICENSE
├── README.md
├── go.mod
└── go.sum
```

## 🛠 Technologies

- Go
- GitHub REST API
- Standard Library (`net/http`, `encoding/json`)

## 📖 API

This project uses the following GitHub REST API endpoint:

```
GET https://api.github.com/users/{username}/events
```

## 🧪 Run

```bash
go run ./cmd/github-activity octocat
```

## 📄 License

This project is licensed under the MIT License.

---

Project idea from the roadmap.sh Backend Developer roadmap.