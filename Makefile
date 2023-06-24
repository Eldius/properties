
test:
	go test ./... -cover

vulncheck:
	govulncheck ./...

lint:
	golangci-lint run ./...
