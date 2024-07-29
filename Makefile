generate:
	go generate ./example/...

test: generate
	go test -count=1 -v ./example

.PHONY: generate test
