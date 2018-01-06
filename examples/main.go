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
	ID        int    `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

type Empty struct{}

func main() {
	repo := Repository{
		ID:          123,
		Name:        "ltsv-go",
		Description: "ltsv Marshal and Unmarshal",
		URL:         "http://google.com",
		Topics:      []string{"go", "ltsv", "reflect"},
		License:     map[string]string{"key": "value", "hoge": "foo"},
		If:          []string{"interface"},
		Owner:       Owner{ID: 999, Login: "naoto0822"},
	}

	ret := ltsv.Marshal(repo)
	fmt.Println("Marshal result: \n", ret)

	target := Repository{}
	err := ltsv.Unmarshal(ret, &target)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("Unmarshal result: \n", target)
}
