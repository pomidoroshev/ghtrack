package main

import (
	"fmt"
	"time"
)

func main() {
	options := ParseOptions()
	config := NewConfig(configFileName)
	extractor := NewTimespanExtractor(timespanPattern)

	loader := NewGitHubLoader(options.Owner, options.Repository, options.Milestone, options.IssueState, config.Token)

	issues := make(chan Measurable, 1)
	go loader.Load(issues)

	var totalTime time.Duration
	for issue := range issues {
		fmt.Printf("%s - %s\n", issue.Title(), issue.Elapsed(extractor))
		totalTime += issue.Elapsed(extractor)
	}

	fmt.Printf("Total: %s", totalTime)
}
