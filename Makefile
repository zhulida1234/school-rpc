GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

school-rpc:
	./bin/compile.sh
	env GO111MODULE=on go build -v -o school-rpc $(LDFLAGS) ./cmd

clean:
	rm school-rpc

test:
	go test -v ./...

lint:
	golangci-lint run ./...