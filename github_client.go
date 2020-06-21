package main

import (
	"context"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	ctx    context.Context
	client *github.Client
}

func NewGitHubClient(accessToken string) GitHubClient {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return GitHubClient{
		ctx:    ctx,
		client: client,
	}
}

func (c GitHubClient) Issues(owner, repository, issueState, milestone string, ch chan Measurable) {
	opt := &github.IssueListByRepoOptions{
		State:       issueState,
		Milestone:   milestone,
		ListOptions: github.ListOptions{PerPage: perPage},
	}
	for {
		issues, resp, err := c.client.Issues.ListByRepo(c.ctx, owner, repository, opt)
		if err != nil {
			log.Fatalf("Error fetching issues: %v", err)
		}
		for _, issue := range issues {
			comments := c.Comments(owner, repository, issue.GetNumber())
			ch <- Issue{
				title:    issue.GetTitle(),
				Comments: comments,
			}
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
		break
	}
	close(ch)
}

func (c GitHubClient) Comments(owner, repository string, issueNumber int) []Comment {
	var result []Comment
	opt := &github.IssueListCommentsOptions{
		ListOptions: github.ListOptions{PerPage: perPage},
	}
	for {
		comments, resp, err := c.client.Issues.ListComments(
			c.ctx, owner, repository, issueNumber, opt,
		)
		if err != nil {
			log.Fatalf("Error fetching comments: %v", err)
		}
		for _, comment := range comments {
			result = append(result, Comment{text: comment.GetBody()})
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return result
}
