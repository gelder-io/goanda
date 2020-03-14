# Linter variables.
LINTVER="v1.23.6"
LINTSRC="https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh"


# Linting and code checks.
.PHONY: fmt outdated lint
fmt:
	goimports -w .
	go mod tidy

lint: install-linter
	golangci-lint run

outdated:
	go list -u -m -json all | docker run -i psampaz/go-mod-outdated -update -direct -ci


# Environment setup
.PHONY: install-linter
install-linter:
	which golangci-lint || curl -sfL $(LINTSRC) | sh -s -- -b $(shell go env GOPATH)/bin $(LINTVER)
