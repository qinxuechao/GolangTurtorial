package main

import (
	"fmt"
	"time"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"log"
	"os"
)

const IssueURL = "https://developer.github.com/v3/search/#search-issues."

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"created_at"`
	Body     string    //Markdown 格式
}

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// SearchIssues 函数查询 GitHub 的 issue 跟踪接口
func SearchIssues(terms [] string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// 我们必须在所有可能的分支上关闭resp.Body
	// 第五章将讲述defer, 它可以让代码更简单一点
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%5-d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

