SRC = *.go
DIST = bin/main

default: $(DIST)

$(DIST): $(SRC)
	go build -o $(DIST) $(SRC)

.PHONY: clean
clean:
	rm $(DIST)

.PHONY: run
run:
	go run .

.PHONY: test
test:
	go test

.PHONY: cov
cov:
	go test -coverprofile=.coverage
	go tool cover -html=.coverage

.PHONY: check
check:
	golint ./...

.PHONY: format
format:
	gofmt -s -w .

.PHONY: imports
imports:
	goimports -w .

.PHONY: fmt
fmt: format imports
