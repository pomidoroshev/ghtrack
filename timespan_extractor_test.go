package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTimespanExtractor(t *testing.T) {
	assert.NotEmpty(t, NewTimespanExtractor(`#t ([\d\w]+)`))
}

func TestNewTimespanExtractorPanic(t *testing.T) {
	assert.Panics(t, func() { NewTimespanExtractor(`#t ([\d\+)`) })
}

func TestNewTimespanExtractorParse(t *testing.T) {
	assert := assert.New(t)
	elapser := NewTimespanExtractor(`#t ([\d\w]+)`)
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"#t", 0},
		{"#t 1h", 3600},
		{"#t 1h30m", 5400},
		{"#t 1h #t 1h", 7200},
	}
	for _, test := range tests {
		seconds, err := elapser.Parse(test.input)
		assert.Equal(test.expected, seconds)
		assert.Nil(err)
	}
}

func TestNewTimespanExtractorParseError(t *testing.T) {
	assert := assert.New(t)
	elapser := NewTimespanExtractor(`#t ([\d\w]+)`)
	inputs := []string{
		"#t 1",
		"#t h1",
		"#t h",
		"#t 1hhhh",
	}
	for _, input := range inputs {
		seconds, err := elapser.Parse(input)
		assert.Equal(0, seconds)
		assert.Error(err)
	}
}
