NAME = ghtrack
SRC = ./app
BINDIR = ./bin
DIST = $(BINDIR)/$(NAME)
DIST_LINUX_AMD64 = $(DIST).linux-amd64
DIST_LINUX_386 = $(DIST).linux-386
DIST_LINUX_ARM = $(DIST).linux-arm
DIST_LINUX_ARM64 = $(DIST).linux-arm64
DIST_WINDOWS_AMD64 = $(DIST).windows-amd64.exe
DIST_DARWIN_AMD64 = $(DIST).darwin-amd64
DIST_FREEBSD_AMD64 = $(DIST).freebsd-amd64

DISTS = $(DIST) $(DIST_LINUX_AMD64) $(DIST_LINUX_386) $(DIST_LINUX_ARM) \
		$(DIST_LINUX_ARM64) $(DIST_WINDOWS_AMD64) $(DIST_DARWIN_AMD64) \
		$(DIST_FREEBSD_AMD64)

BUILD = go build

default: build

build:
	$(BUILD) -o $(DIST) $(SRC)

release:
	$(BUILD) -ldflags="-s -w" -o $(DIST) $(SRC)

build-archs:
	GOOS=linux GOARCH=amd64 $(BUILD) -o $(DIST_LINUX_AMD64) -ldflags "-s -w" $(SRC)
	GOOS=linux GOARCH=386 $(BUILD) -o $(DIST_LINUX_386) -ldflags "-s -w" $(SRC)
	GOOS=linux GOARCH=arm $(BUILD) -o $(DIST_LINUX_ARM) -ldflags "-s -w" $(SRC)
	GOOS=linux GOARCH=arm64 $(BUILD) -o $(DIST_LINUX_ARM64) -ldflags "-s -w" $(SRC)
	GOOS=windows GOARCH=amd64 $(BUILD) -o $(DIST_WINDOWS_AMD64) -ldflags "-s -w" $(SRC)
	GOOS=darwin GOARCH=amd64 $(BUILD) -o $(DIST_DARWIN_AMD64) -ldflags "-s -w" $(SRC)
	GOOS=freebsd GOARCH=amd64 $(BUILD) -o $(DIST_FREEBSD_AMD64) -ldflags "-s -w" $(SRC)

.PHONY: clean
clean:
	$(foreach dist,$(DISTS),rm -f $(dist);)

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
