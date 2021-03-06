SSHKEY=""
DEPLOY_HOST="deploy@brigade.ec2.shopify.com"

BRANCH=`git rev-parse --abbrev-ref HEAD`
COMMIT=`git rev-parse --short HEAD`
GOLDFLAGS="-X main.branch $(BRANCH) -X main.commit $(COMMIT)"

all: test build

setup:
	@go get -u "github.com/tools/godep"
	@go get -u "github.com/golang/lint/golint"
	@go get -u "github.com/kisielk/errcheck"

# http://cloc.sourceforge.net/
cloc:
	@cloc --sdir='Godeps' --not-match-f='Makefile|_test.go' .

errcheck:
	@echo "=== errcheck ==="
	@errcheck github.com/Shopify/brigade/...

vet:
	@echo "==== go vet ==="
	@go vet ./...

lint:
	@echo "==== go lint ==="
	@golint ./**/*.go

fmt:
	@echo "=== go fmt ==="
	@go fmt ./...

install: test
	@echo "=== go install ==="
	@godep go install -ldflags=$(GOLDFLAGS)

build:
	@echo "=== go build ==="
	@mkdir -p bin/
	@godep go build -ldflags=$(GOLDFLAGS) -o bin/brigade

test: fmt vet lint errcheck
	@echo "=== go test ==="
	@godep go test ./... -cover

deploy: test
	# Compile
	@mkdir -p bin/
	GOARCH=amd64 GOOS=linux godep go build -ldflags=$(GOLDFLAGS) -o bin/brigade
	# Copy binaries
	@scp bin/brigade $(DEPLOY_HOST):~/
	# Cleanup binaries
	@rm bin/brigade

.PHONY: setup cloc errcheck vet lint fmt install build test deploy
