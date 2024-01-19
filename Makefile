.PHONY: test

dev:
	nodemon --exec go run cmd/main.go --signal SIGTERM

test:
	go test ./internal/... -coverprofile=coverage.out
	
cover:
	go tool cover -html=coverage.out