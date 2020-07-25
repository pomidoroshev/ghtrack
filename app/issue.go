package main

import "time"

type Measurable interface {
	Title() string
	Elapsed(extractor TimespanExtractor) time.Duration
}

// Issue represents project issue
type Issue struct {
	title    string
	Comments []Comment
}

// NewIssue creates an issue with given comments
func NewIssue(title string, comments []Comment) Issue {
	return Issue{title: title, Comments: comments}
}

// Elapsed returns total durations from all comments
func (i Issue) Elapsed(extractor TimespanExtractor) time.Duration {
	var duration time.Duration
	for _, c := range i.Comments {
		duration += c.Elapsed(extractor)
	}
	return duration
}

// Title returns issue title
func (i Issue) Title() string {
	return i.title
}

// Comment is an issue comment that may contain timespan
type Comment struct {
	text string
}

// NewComment creates a comment with text and elapser
func NewComment(text string) Comment {
	return Comment{text: text}
}

// Elapsed returns duration from parsed timespan in comment text
func (c Comment) Elapsed(extractor TimespanExtractor) time.Duration {
	duration, err := extractor.Parse(c.text)
	if err != nil {
		return 0
	}
	return duration
}
