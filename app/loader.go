package main

// Loader fetches issues from different sources
type Loader interface {
	Load(ch chan Measurable)
}

// GitHubLoader loads issues from GitHub
type GitHubLoader struct {
	owner      string
	repository string
	milestone  string
	issueState string

	client GitHubClient
}

// NewGitHubLoader creates a Loader object with GitHub client.
func NewGitHubLoader(owner, repository, milestone, issueState, accessToken string) GitHubLoader {
	client := NewGitHubClient(accessToken)
	return GitHubLoader{
		owner:      owner,
		repository: repository,
		milestone:  milestone,
		issueState: issueState,
		client:     client,
	}
}

// Load fetches repository issues with comments and sends them to channel.
// Sends nil when it finishes.
func (l GitHubLoader) Load(ch chan Measurable) {
	l.client.Issues(l.owner, l.repository, l.issueState, l.milestone, ch)
}
