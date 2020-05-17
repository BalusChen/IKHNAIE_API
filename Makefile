
build:
	go build -o ./output/ikhnaie
run:
	go run main.go

fmt:
	go fmt ./...

tidy:
	go mod tidy

lint:
ifeq (, $(shell which golangci-lint))
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
endif
	$(GOPATH)/bin/golangci-lint run
