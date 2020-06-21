package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	All = iota
	Open
	Closed
)

var issueState = map[int]string{
	All:    "all",
	Open:   "open",
	Closed: "closed",
}

type Options struct {
	Owner      string
	Repository string
	IssueState string
	Milestone  string
}

func ParseOptions() Options {
	flag.Usage = usage

	issueStatePtr := flag.String("s", issueState[All], "Issue state: all, open, closed")
	milestonePtr := flag.String("m", "*", "Milestone number: milestone number, *, none")

	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		usage()
		os.Exit(1)
	}

	return Options{
		Owner:      args[0],
		Repository: args[1],
		IssueState: *issueStatePtr,
		Milestone:  *milestonePtr,
	}
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] owner repository\n", os.Args[0])
	flag.PrintDefaults()
}
