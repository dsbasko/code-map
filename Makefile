.PHONY: build install-deps lint test test-cover test-cover-svg test-cover-html infra-start infra-stop
.SILENT:

build:
	@echo "Building the binaries..."
	@GOOS=linux GOARCH=amd64 go build -o $(CURDIR)/bin/linux_amd64
	@GOOS=linux GOARCH=arm64 go build -o $(CURDIR)/bin/linux_arm64
	@GOOS=darwin GOARCH=amd64 go build -o $(CURDIR)/bin/darwin_amd64
	@GOOS=darwin GOARCH=arm64 go build -o $(CURDIR)/bin/darwin_arm64
	@GOOS=windows GOARCH=amd64 go build -o $(CURDIR)/bin/windows_amd64.exe
	@GOOS=windows GOARCH=386 go build -o $(CURDIR)/bin/windows_386.exe



# Dependencies
install-deps:
	@GOBIN=$(CURDIR)/temp/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go mod tidy



# Lint
lint:
	@$(CURDIR)/temp/bin/golangci-lint run -c .golangci.yml --path-prefix . --fix



# Infrastructure
infra-start: infra-stop
	@docker compose \
		-f ./infrastructure/docker-compose.yaml \
		up --build -d;
infra-stop:
	@docker compose \
		-f ./infrastructure/docker-compose.yaml \
		 down > /dev/null 2>&1 || true