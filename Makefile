all: dep vet lint test

dep:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/golang/lint/golint

vet:
	go vet ./ltsv

lint:
	golint -set_exit_status ./ltsv

test:
	go test -v ./ltsv

.PHONY: dep vet test lint
