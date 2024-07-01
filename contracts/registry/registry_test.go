package registry_test

import (
	libtesting "github.com/pyropy/registry/lib/testing"
	"math/big"
	"testing"
)

func TestFileRegistry(t *testing.T) {

	const numAccounts = 1
	const autoCommit = true
	var accountBalance = big.NewInt(100)

	backend, err := libtesting.NewSimulatedBackend(numAccounts, autoCommit, accountBalance)

	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}
	defer backend.Close()

	client, err := libtesting.NewClient(backend, backend.PrivateKeys[0])
	if err != nil {
		t.Fatalf("unable to create ethereum api: %s", err)
	}

}
