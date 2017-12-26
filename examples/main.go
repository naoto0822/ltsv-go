package main

import (
	// "fmt"
	"github.com/naoto0822/ltsv-go/ltsv"
)

type Repository struct {
	ID          int                    `ltsv:"id"`
	Name        string                 `ltsv:"name"`
	Description string                 `ltsv:"description"`
	URL         string                 `ltsv:"url"`
	Private     bool                   `ltsv:"private"`
	Topics      []string               `ltsv:"topics"`
	Owner       map[string]interface{} `ltsv:"owner"`
}

func main() {
	repo := Repository{
		ID:          123,
		Name:        "ltsv-go",
		Description: "ltsv Marshal and Unmarshal",
		URL:         "http://google.com",
	}

	ltsv.Marshal(repo)
}
