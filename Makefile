all:
	go generate ./pkg/...
	go generate ./example/...

test: all
	go test -count=1 -v ./example

.PHONY: all test
