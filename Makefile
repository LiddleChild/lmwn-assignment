.PHONY: test

dev:
	. ./export-env.sh ; nodemon --exec go run cmd/main.go --signal SIGTERM

run:
	. ./export-env.sh ; go run cmd/main.go

test:
	go test ./internal/... -coverprofile=coverage.out
	
cover:
	go tool cover -html=coverage.out