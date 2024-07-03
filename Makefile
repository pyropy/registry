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

#? compile: Compile the solidity contract.
compile:
	@solc --optimize --bin --abi contract/src/FileRegistry.sol -o contract/abi --overwrite
	@abigen --bin contract/abi/FileRegistry.bin --abi contract/abi/FileRegistry.abi --pkg contract --type FileRegistry --out contract/contract.go


#? deploy: Deploy the solidity contract.
deploy:
	@go run cmd/deploy/main.go

#? test: Run the tests.
test:
	@echo "Testing..."
	@go test ./... -v

#? clean: Remove all executables and keys.
clean:
	@echo "Cleaning..."
	@rm -rf bin
	@rm -rf datadir

#? lint: Lint whole repository.
lint: install-golangcilint
	$(GOLANGCI_LINT) run ./...

#? install-golangcilint: Install golangci-lint.
install-golangcilint:
	test -f $(GOLANGCI_LINT) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$($(GO) env GOPATH)/bin $(GOLANGCI_LINT_VERSION)

#? devtools: Install recommended developer tools.
devtools:
	brew update
	brew list ipfs || brew install ipfs
	brew list ethereum || brew install ethereum
	brew list solidity || brew install solidity


#? devtools-update: Update recommended developer tools.
devtools-update:
	brew update
	brew list ipfs || brew upgrade ipfs
	brew list ethereum || brew upgrade ethereum
	brew list solidity || brew upgrade solidity


#? dev-keygen: Generate a new development key (WARNING! Do not use for production).
dev-keygen:
	scripts/dev-keygen.sh

#? help: Get more info on make commands.
help: Makefile
	@echo ''
	@echo 'Usage:'
	@echo '  make [target]'
	@echo ''
	@echo 'Targets:'
	@sed -n 's/^#?//p' $< | column -t -s ':' |  sort | sed -e 's/^/ /'