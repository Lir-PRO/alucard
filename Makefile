test-config:
	@echo "Running tests for internal/config..."
	@go test -v ./internal/config

cli-start:
	@echo "Running CLI application directly..."
	go run ./internal/cli/cli.go
