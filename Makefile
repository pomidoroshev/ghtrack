SRC = *.go
DIST = bin/main

default: $(DIST)

$(DIST): $(SRC)
	go build -o $(DIST) $(SRC)

.PHONY: clean
build-alpine:
	CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o bin/main .

.PHONY: clean
clean:
	rm $(DIST)

.PHONY: run
run:
	go run .

.PHONY: test
test:
	go test

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
