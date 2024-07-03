package registry_test

import (
	"context"
	"github.com/pyropy/eth"
	"github.com/pyropy/registry/config"
	"github.com/pyropy/registry/contract"
	"github.com/pyropy/registry/ipfs"
	"github.com/pyropy/registry/registry"
	"math/big"
	"testing"
)

const (
	numAccounts = 1
	autoCommit  = true
	gasLimit    = 1600000
)

var accountBalance = big.NewInt(100)

func TestRegistry_UploadFile(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// create simulated backend and connected client
	backend, err := eth.CreateSimulatedBackend(numAccounts, autoCommit, accountBalance)
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}

	defer backend.Close()

	client, err := eth.NewClient(backend, backend.PrivateKeys[0])
	if err != nil {
		t.Fatalf("unable to create ethereum api: %s", err)
	}

	address, _, err := contract.DeployContract(ctx, client, gasLimit)
	if err != nil {
		t.Fatalf("unable to create env: %s", err)
	}

	cfg := &config.Config{
		EthGasLimit:                 gasLimit,
		EthPrivateKey:               backend.PrivateKeys[0],
		FileRegistryContractAddress: address,
	}

	ipfsClient := ipfs.NewMockClient()
	reg, err := registry.NewRegistry(cfg, ipfsClient, backend)
	if err != nil {
		t.Fatalf("unable to create registry: %s", err)
	}

	want := "bafybeigdyrzt5sfp7udm7hu76uh7y26nf3efuylqabf3oclgtqy55fbzdi"
	got, err := reg.UploadFile(ctx, "test", nil)
	if err != nil {
		t.Fatalf("unable to upload file: %s", err)
	}

	if got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}
