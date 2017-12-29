all: dep vet lint test

dep:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/golang/lint/golint

vet:
	go tool vet ./ltsv

test:
	go test -v ./ltsv

lint:
	golint -set_exit_status ./ltsv

.PHONY: dep vet test lint
