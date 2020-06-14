package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentElapsed(t *testing.T) {
	assert := assert.New(t)
	extractor := NewTimespanExtractor(`#t ([\d\w]+)`)
	tests := []struct {
		text    string
		seconds int
	}{
		{"Hello #t 1h", 3600},
	}

	for _, test := range tests {
		comment := NewComment(test.text, &extractor)
		assert.Equal(test.seconds, comment.Elapsed())
	}
}

func TestIssueElapsed(t *testing.T) {
	assert := assert.New(t)
	extractor := NewTimespanExtractor(`#t ([\d\w]+)`)
	tests := []struct {
		comments *[]Comment
		seconds  int
	}{
		{&[]Comment{NewComment("#t 1h", &extractor), NewComment("#t 30m", &extractor)}, 5400},
		{&[]Comment{NewComment("#t 1", &extractor)}, 0},
		{&[]Comment{}, 0},
		{nil, 0},
	}

	for _, test := range tests {
		issue := NewIssue(test.comments)
		assert.Equal(test.seconds, issue.Elapsed())
	}
}
