package contract

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pyropy/eth"
	"math/big"
)

var (
	zeroGwei = big.NewFloat(0.0)
)

// DeployContract deploys FileRegistry contract with given client and gas limit
func DeployContract(ctx context.Context, client *eth.Client, gasLimit uint64) (common.Address, *FileRegistry, error) {
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("unable to suggest gas price: %s", err)
	}

	tranOpts, err := client.NewTransactOpts(ctx, gasLimit, gasPrice, zeroGwei)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("unable to create transaction opts for deploy: %s", err)
	}

	contractAddress, tx, _, err := DeployFileRegistry(tranOpts, client.Backend)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("unable to deploy FileRegistry: %s", err)
	}

	if _, err := client.WaitMined(ctx, tx); err != nil {
		return common.Address{}, nil, fmt.Errorf("waiting for deploy: %s", err)
	}

	testRegistry, err := NewFileRegistry(contractAddress, client.Backend)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("error creating FileRegistry: %s", err)
	}

	return contractAddress, testRegistry, err
}
