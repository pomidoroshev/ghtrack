package main

import (
	"regexp"
	"time"
)

// TimespanExtractor parses timespan from string and returns duration
type TimespanExtractor struct {
	pattern *regexp.Regexp
}

// NewTimespanExtractor creates new extractor instance with given regexp pattern
func NewTimespanExtractor(pattern string) TimespanExtractor {
	return TimespanExtractor{regexp.MustCompile(pattern)}
}

// Parse returns number of duration of all timespans in given text
func (e *TimespanExtractor) Parse(text string) (time.Duration, error) {
	var total time.Duration
	matches := e.pattern.FindAllStringSubmatch(text, -1)
	for _, groups := range matches {
		duration, err := time.ParseDuration(groups[1])
		if err != nil {
			return 0, err
		}
		total += duration
	}
	return total, nil
}
