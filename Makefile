.PHONY: build
build:
	go build cmd/bss-company-api/main.go

.PHONY: test
test:
	go test -v ./...