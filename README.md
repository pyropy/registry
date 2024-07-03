# File Registry

![robot](banner.jpg)

File registry is a simple application that allows you to upload files to IPFS and register them on any EVM compatible blockchain.

## Usage

To see the list of available commands, run the `make help` command:

```bash
Usage:
  make [target]

Targets:
  all                    Build all packages and executables.
  build                  Build the app.
  clean                  Remove all executables and keys.
  clean                  Run the app.
  compile                Compile the solidity contract.
  deploy                 Deploy the solidity contract.
  dev-keygen             Generate a new development key (WARNING! Do not use for production).
  devtools               Install recommended developer tools.
  devtools-update        Update recommended developer tools.
  help                   Get more info on make commands.
  install-golangcilint   Install golangci-lint.
  lint                   Lint whole repository.
  test                   Run the tests.

```


## Dependencies

File registry requires the following dependencies to be installed:

| Dependency | Usage                                                                                    |
|:----------:|------------------------------------------------------------------------------------------|
| **`geth`** | Used to generate private key used for signing the transactions.                  |
|   `solc`   | Used to compile Solidity contract to bytecode and generate JSON ABI file.        |
|  `abigen`  | Used to generate Go interface for deploying and interacting with the contract. |

### Installing dependencies

To easily install the dependencies using `brew`, run the following command:

```bash
make devtools
```


## Building from the source

In order to build file registry binary, run the following command:

```bash
make build
```

If you wish to compile the Solidity contract before building the binary, you can run the following command:

```bash
make all
```


## Generating private key
Before running the application, you need to generate a private key that will be used for signing the transactions. To generate a new private key, run the following command:

```bash
make dev-keygen
```

This will generate a new private key saved under `/tmp/eth_pk_file` and password file under `/tmp/eth_pk_password`.

**WARNING!** Do not use the generated private key for production purposes. It is recommended to use a hardware wallet or a secure key management system for production.

## Running the application

In order to run the application you will need to setup your environment variables first. Example `.env.example` file is provided in the repository. You can copy the example file and adjust the values to your needs.

Once you have configured the environment variables, you can run the application by running previously built binary or by executing `make run` command.


## Example requests

### Upload file and register it in the file registry

Example request:
```bash
curl --location 'localhost:8000/v1/files' \
--form 'filePath="/tmp/coin.png"' \
--form 'file=@"/path/to/file"'
```

Example response:
```json
{
  "cid": "/ipfs/bafkreihxhjf56iyjmxshsjsd5hjeqzpsmt2ftsltr6ykdat55kemzyo5uu"
}
```

### Retrieve file CID from the file registry

Example request:
```bash
curl --location 'http://localhost:8000/v1/files?filePath=%2Ftmp%2Fcoin.png'
```

Example response:

```json
{
  "cid": "/ipfs/bafkreihxhjf56iyjmxshsjsd5hjeqzpsmt2ftsltr6ykdat55kemzyo5uu"
}
```