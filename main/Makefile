PHONY: build-binary
build-binary: clean fmt
	rm -f go.sum
	go get -v -t -d ./...
	$(eval DIR := $(GOPATH)/src/github.com/rashmi43/go-messenger/main)
	CGO_ENABLED=0 GO111MODULE=on GOOS=$(GOOS) go build -o $(GOPATH)/bin/go-messenger
	go mod tidy
	rm -f go.sum
	
PHONY: fmt
fmt:
	gofmt -w *.go */*.go
	

PHONY: clean
clean:
	rm -f $(GOPATH)/bin/go-messenger
	rm -f go.sum
