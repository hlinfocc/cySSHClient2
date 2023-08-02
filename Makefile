export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w

all: fmt build

build: cyssh cyscp cysshserver cymain


fmt:
	go fmt ./...

fmt-more:
	gofumpt -l -w .

vet:
	go vet ./...

cyssh:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/cyssh ./cmd/cyssh

cyscp:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/cyscp ./cmd/cyscp
	
cysshserver:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/csc-server ./server

cymain:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/cysshmain ./

test: gotest

gotest:
	go test -v --cover ./assets/...
	go test -v --cover ./cmd/...
	go test -v --cover ./client/...
	go test -v --cover ./server/...
	go test -v --cover ./pkg/...

	
clean:
	rm -f ./bin/cyssh
	rm -f ./bin/cyscp
	rm -f ./bin/cyssh-server
