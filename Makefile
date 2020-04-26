
run:
	go run main.go

fmt:
	go fmt ./...

lint:
ifeq (, $(shell which golangci-lint))
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
endif
	$(GOPATH)/bin/golangci-lint run
