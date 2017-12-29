package main

import (
	"fmt"
	"github.com/naoto0822/ltsv-go/ltsv"
)

type Repository struct {
	ID          int               `ltsv:"id"`
	Name        string            `ltsv:"name"`
	Description string            `ltsv:"description"`
	URL         string            `ltsv:"url"`
	Private     bool              `ltsv:"private"`
	Topics      []string          `ltsv:"topics"`
	License     map[string]string `ltsv:"license"`
	If          interface{}       `ltsv:"interface"`
	Owner       Owner             `ltsv:"owner"`
	Empty       Empty             `ltsv:"empty"`
}

type Owner struct {
	ID        int
	Login     string
	AvatarURL string
}

type Empty struct{}

func main() {
	repo := Repository{
		ID:          123,
		Name:        "ltsv-go",
		Description: "ltsv Marshal and Unmarshal",
		URL:         "http://google.com",
		Topics:      []string{"go", "ltsv", "reflect"},
		License:     map[string]string{},
	}

	ret := ltsv.Marshal(repo)
	fmt.Println(ret)
}
