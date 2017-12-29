package main

import (
	"fmt"
	"github.com/naoto0822/ltsv-go/ltsv"
)

type Repository struct {
	ID          int                    `ltsv:"id"`
	Name        string                 `ltsv:"name"`
	Description string                 `ltsv:"description"`
	URL         string                 `ltsv:"url"`
	Private     bool                   `ltsv:"private"`
	Topics      []string               `ltsv:"topics"`
	License     map[string]interface{} `ltsv:"license"`
	Owner       Owner                  `ltsv:owner`
}

type Owner struct {
	ID        int
	Login     string
	AvatarURL string
}

func main() {
	repo := Repository{
		ID:          123,
		Name:        "ltsv-go",
		Description: "ltsv Marshal and Unmarshal",
		URL:         "http://google.com",
		Topics:      []string{"go", "ltsv", "reflect"},
	}

	ret, err := ltsv.Marshal(repo)
	fmt.Println(ret)
	fmt.Println(err)
}
