clean:
	@rm -fr ./cover.out

lint-setup:
	@go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest

lint: lint-setup
	@golangci-lint run ./...

lint-fix: lint-setup
	@golangci-lint run --fix ./...

test-setup:
	@go install gotest.tools/gotestsum@latest

test: test-setup
	@gotestsum -- -timeout 300s ./...

test-coverage: test-setup
	@gotestsum -- -timeout 300s -coverprofile=./cover.out ./...

test-coverage-html: test-coverage
	@go tool cover -html=./cover.out

test-watch: test-setup
	@gotestsum --watch ./...
