package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCommentElapsed(t *testing.T) {
	assert := assert.New(t)
	extractor := NewTimespanExtractor(timespanPattern)
	tests := []struct {
		text     string
		duration time.Duration
	}{
		{"Hello #t 1h", 3600},
	}

	for _, test := range tests {
		comment := NewComment(test.text)
		assert.Equal(test.duration*time.Second, comment.Elapsed(extractor))
	}
}

func TestIssueElapsed(t *testing.T) {
	assert := assert.New(t)
	extractor := NewTimespanExtractor(timespanPattern)
	tests := []struct {
		comments []Comment
		duration time.Duration
	}{
		{[]Comment{NewComment("#t 1h"), NewComment("#t 30m")}, 5400},
		{[]Comment{NewComment("#t 1")}, 0},
		{[]Comment{}, 0},
	}

	for _, test := range tests {
		issue := NewIssue("some title", test.comments)
		assert.Equal(test.duration*time.Second, issue.Elapsed(extractor))
	}
}
