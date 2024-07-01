.PHONY: install-deps
install-deps:
	@npm install -g solc
	@go install github.com/ethereum/go-ethereum/cmd/abigen@latest

.PHONY: compile
compile:
	@solc --abi contracts/registry/src/FileRegistry.sol -o contracts/registry/abi --overwrite

.PHONY: gen
gen: compile
	@abigen --abi contracts/registry/abi/FileRegistry.abi --pkg registry --type FileRegistry --out contracts/registry/registry.go

