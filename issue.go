package main

// Trackable object can be measured in seconds
type Trackable interface {
	Elapsed() int
}

// Issue represents project issue
type Issue struct {
	comments *[]Comment
}

// NewIssue creates an issue with given comments
func NewIssue(comments *[]Comment) Issue {
	return Issue{comments: comments}
}

// Elapsed returns sum of seconds from all comments
func (i *Issue) Elapsed() int {
	if i.comments == nil {
		return 0
	}
	seconds := 0
	for _, c := range *i.comments {
		seconds += c.Elapsed()
	}
	return seconds
}

// Comment is an issue comment that may contain timespan
type Comment struct {
	text      string
	extractor *TimespanExtractor
}

// NewComment creates a comment with text and elapser
func NewComment(text string, extractor *TimespanExtractor) Comment {
	return Comment{
		text:      text,
		extractor: extractor,
	}
}

// Elapsed returns seconds from parsed timespan in comment text
func (c *Comment) Elapsed() int {
	seconds, err := c.extractor.Parse(c.text)
	if err != nil {
		return 0
	}
	return seconds
}
