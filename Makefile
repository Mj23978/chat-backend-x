.PHONY: lint
lint:
		GO111MODULE=on golangci-lint run -v ./...
