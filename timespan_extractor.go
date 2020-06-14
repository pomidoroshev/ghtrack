package main

import (
	"regexp"
	"time"
)

// TimespanExtractor parses timespan from string and returns seconds
type TimespanExtractor struct {
	pattern *regexp.Regexp
}

// NewTimespanExtractor creates new extractor instance with given regexp pattern
func NewTimespanExtractor(pattern string) TimespanExtractor {
	return TimespanExtractor{regexp.MustCompile(pattern)}
}

// Parse returns number of seconds of all timespans in given text
func (e *TimespanExtractor) Parse(text string) (int, error) {
	seconds := 0
	matches := e.pattern.FindAllStringSubmatch(text, -1)
	for _, groups := range matches {
		duration, err := time.ParseDuration(groups[1])
		if err != nil {
			return 0, err
		}
		seconds += int(duration.Seconds())
	}
	return seconds, nil
}
