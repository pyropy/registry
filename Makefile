GO ?= go
GOBIN ?= $$($(GO) env GOPATH)/bin
GOLANGCI_LINT ?= $(GOBIN)/golangci-lint
GOLANGCI_LINT_VERSION ?= v1.56.2

.PHONY: all build run compile test clean lint install-golangcilint devtools devtools-update help

#? all: Build all packages and executables.
all: compile build

#? build: Build the app.
build:
	@echo "Building..."
	
	@go build -o bin/main cmd/app/main.go

#? clean: Run the app.
run:
	@go run cmd/app/main.go

#? compile: Compile solidity contract.
compile:
	@solc --bin --abi contract/src/FileRegistry.sol -o contract/abi --overwrite
	@abigen --bin contract/abi/FileRegistry.bin --abi contract/abi/FileRegistry.abi --pkg contract --type FileRegistry --out contract/contract.go

#? test: Run the tests.
test:
	@echo "Testing..."
	@go test ./... -v

#? clean: Remove all executables.
clean:
	@echo "Cleaning..."
	@rm -rf bin

#? lint: Lint whole repository.
lint: install-golangcilint
	$(GOLANGCI_LINT) run ./...

#? install-golangcilint: Install golangci-lint.
install-golangcilint:
	test -f $(GOLANGCI_LINT) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$($(GO) env GOPATH)/bin $(GOLANGCI_LINT_VERSION)

#? devtools: Install recommended developer tools.
devtools:
	brew update
	brew list ethereum || brew install ethereum
	brew list solidity || brew install solidity


#? devtools-update: Update recommended developer tools.
devtools-update:
	brew update
	brew list ethereum || brew upgrade ethereum
	brew list solidity || brew upgrade solidity


#? help: Get more info on make commands.
help: Makefile
	@echo ''
	@echo 'Usage:'
	@echo '  make [target]'
	@echo ''
	@echo 'Targets:'
	@sed -n 's/^#?//p' $< | column -t -s ':' |  sort | sed -e 's/^/ /'