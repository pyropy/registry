package contract_test

import (
	"context"
	"math/big"
	"testing"

	eth "github.com/pyropy/eth"
	"github.com/pyropy/eth/currency"
	"github.com/pyropy/registry/contract"
)

func TestFileRegistry(t *testing.T) {
	ctx := context.Background()

	const numAccounts = 1
	const autoCommit = true
	var accountBalance = big.NewInt(100)

	backend, err := eth.CreateSimulatedBackend(numAccounts, autoCommit, accountBalance)
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}
	defer backend.Close()

	client, err := eth.NewClient(backend, backend.PrivateKeys[0])
	if err != nil {
		t.Fatalf("unable to create ethereum api: %s", err)
	}

	// =========================================================================

	const gasLimit = 1600000
	valueGwei := big.NewFloat(0.0)
	gasPrice := currency.GWei2Wei(big.NewFloat(39.576))
	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, gasPrice, valueGwei)
	if err != nil {
		t.Fatalf("unable to create transaction opts for deploy: %s", err)
	}

	contractID, tx, _, err := contract.DeployFileRegistry(tranOpts, client.Backend)
	if err != nil {
		t.Fatalf("unable to deploy Registry: %s", err)
	}

	if _, err := client.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for deploy: %s", err)
	}

	testRegistry, err := contract.NewFileRegistry(contractID, client.Backend)
	if err != nil {
		t.Fatalf("error creating Registry: %s", err)
	}

	// =========================================================================

	tranOpts, err = client.NewTransactOpts(ctx, gasLimit, gasPrice, valueGwei)
	if err != nil {
		t.Fatalf("unable to create transaction opts for setitem: %s", err)
	}

	key := "/tmp/example.zip"
	value := "QmPK1s3pNYLi9ERiq3BDxKa4XosgWwFRQUydHUtz4YgpqB"

	tx, err = testRegistry.Save(tranOpts, key, value)
	if err != nil {
		t.Fatalf("should be able to set item: %s", err)
	}

	if _, err := client.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for setitem: %s", err)
	}

	// =========================================================================

	callOpts, err := client.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unable to create call opts: %s", err)
	}

	item, err := testRegistry.Items(callOpts, key)
	if err != nil {
		t.Fatalf("should be able to retrieve item: %s", err)
	}

	if item != value {
		t.Fatalf("wrong value, got %s  exp %s", item, value)
	}
}
