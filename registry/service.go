package registry

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pyropy/eth"
	"github.com/pyropy/registry/config"
	"github.com/pyropy/registry/contract"
	"github.com/pyropy/registry/ipfs"
	"io"
	"log"
	"math/big"
)

// TODO: Add logging
type Registry struct {
	ipfsClient       *ipfs.Client
	ethClient        *eth.Client
	registryContract *contract.FileRegistry
}

func NewRegistry(cfg *config.Config, backend eth.Backend) (*Registry, error) {
	ipfsClient, err := ipfs.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	client, err := eth.NewClient(backend, cfg.EthPrivateKey)
	if err != nil {
		return nil, err
	}

	c, err := contract.NewFileRegistry(cfg.FileRegistryContractAddress, client.Backend)
	if err != nil {
		return nil, err
	}

	return &Registry{
		ipfsClient:       ipfsClient,
		ethClient:        client,
		registryContract: c,
	}, nil
}

func (r *Registry) UploadFile(ctx context.Context, filePath string, file io.Reader) (string, error) {
	cid, err := r.ipfsClient.UploadFile(ctx, file)
	if err != nil {
		return "", err
	}

	txOpts, err := r.createTransactionOptions(ctx)
	if err != nil {
		return "", err
	}

	tx, err := r.registryContract.Save(txOpts, filePath, cid)
	if err != nil {
		return "", err
	}

	log.Println("tx hash: ", tx.Hash().Hex())

	return cid, nil
}

func (r *Registry) GetFileCid(ctx context.Context, filePath string) (string, error) {
	callOpts, err := r.ethClient.NewCallOpts(ctx)
	if err != nil {
		return "", err
	}

	return r.registryContract.Get(callOpts, filePath)
}

func (r *Registry) createTransactionOptions(ctx context.Context) (*bind.TransactOpts, error) {
	const gasLimit = 148416
	valueGwei := big.NewFloat(0.0)
	gasPrice, err := r.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	return r.ethClient.NewTransactOpts(ctx, gasLimit, gasPrice, valueGwei)
}
