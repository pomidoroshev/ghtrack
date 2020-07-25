NAME = ghtrack
SRC = ./app
BINDIR = ./bin
DIST = $(BINDIR)/$(NAME)
BUILD = go build

default: $(DIST)

$(DIST): $(SRC)
	$(BUILD) -o $(DIST) $(SRC)

release:
	$(BUILD) -ldflags="-s -w" -o $(DIST) $(SRC)

build-archs:
	GOOS=linux GOARCH=amd64 $(BUILD) -o $(DIST).linux-amd64 -ldflags "-s -w" $(SRC)
	GOOS=linux GOARCH=386 $(BUILD) -o $(DIST).linux-386 -ldflags "-s -w" $(SRC)
	GOOS=linux GOARCH=arm $(BUILD) -o $(DIST).linux-arm -ldflags "-s -w" $(SRC)
	GOOS=linux GOARCH=arm64 $(BUILD) -o $(DIST).linux-arm64 -ldflags "-s -w" $(SRC)
	GOOS=windows GOARCH=amd64 $(BUILD) -o $(DIST).windows-amd64.exe -ldflags "-s -w" $(SRC)
	GOOS=darwin GOARCH=amd64 $(BUILD) -o $(DIST).darwin-amd64 -ldflags "-s -w" $(SRC)
	GOOS=freebsd GOARCH=amd64 $(BUILD) -o $(DIST).freebsd-amd64 -ldflags "-s -w" $(SRC)

.PHONY: clean
clean:
	rm $(DIST)

.PHONY: run
run:
	go run $(SRC)

.PHONY: test
test:
	go test -v $(SRC)

.PHONY: cov
cov:
	go test -coverprofile=.coverage $(SRC)
	go tool cover -html=.coverage

.PHONY: check
check:
	golint $(SRC)/...

.PHONY: format
format:
	gofmt -s -w $(SRC)

.PHONY: imports
imports:
	goimports -w $(SRC)

.PHONY: fmt
fmt: format imports
