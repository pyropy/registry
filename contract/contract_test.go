package contract_test

import (
	"context"
	"github.com/pyropy/eth"
	"github.com/pyropy/registry/contract"
	"math/big"
	"testing"

	"github.com/pyropy/eth/currency"
)

const (
	numAccounts = 1
	autoCommit  = true
	gasLimit    = 1600000
)

var (
	gasPrice       = currency.GWei2Wei(big.NewFloat(39.576))
	valueGwei      = big.NewFloat(0.0)
	accountBalance = big.NewInt(100)
)

func TestFileRegistry(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// create simulated backend and connected client
	backend, err := eth.CreateSimulatedBackend(numAccounts, autoCommit, accountBalance)
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}

	client, err := eth.NewClient(backend, backend.PrivateKeys[0])
	if err != nil {
		t.Fatalf("unable to create ethereum api: %s", err)
	}

	_, testRegistry, err := contract.DeployContract(ctx, client, gasLimit)
	if err != nil {
		t.Fatalf("unable to create env: %s", err)
	}

	// save item
	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, gasPrice, valueGwei)
	if err != nil {
		t.Fatalf("unable to create transaction opts for setitem: %s", err)
	}

	key := "/tmp/example.jpg"
	value := "bafybeigdyrzt5sfp7udm7hu76uh7y26nf3efuylqabf3oclgtqy55fbzdi"

	tx, err := testRegistry.Save(tranOpts, key, value)
	if err != nil {
		t.Fatalf("should be able to set item: %s", err)
	}

	if _, err := client.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for setitem: %s", err)
	}

	// get item
	callOpts, err := client.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unable to create call opts: %s", err)
	}

	item, err := testRegistry.Get(callOpts, key)
	if err != nil {
		t.Fatalf("should be able to retrieve item: %s", err)
	}

	if item != value {
		t.Fatalf("wrong value, got %s  exp %s", item, value)
	}
}
