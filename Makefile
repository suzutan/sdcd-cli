VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT  ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE    ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

LDFLAGS := -X github.com/suzutan/sdcd-cli/cmd.Version=$(VERSION) \
           -X github.com/suzutan/sdcd-cli/cmd.Commit=$(COMMIT) \
           -X github.com/suzutan/sdcd-cli/cmd.BuildDate=$(DATE)

.PHONY: build test lint clean install

build:
	go build -ldflags "$(LDFLAGS)" -o bin/sdcd .

install:
	go install -ldflags "$(LDFLAGS)" .

test:
	go test ./... -v

lint:
	golangci-lint run ./...

clean:
	rm -rf bin/
