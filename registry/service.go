package registry

import (
	"context"
	"fmt"
	"github.com/pyropy/eth"
	"github.com/pyropy/registry/config"
	"github.com/pyropy/registry/contract"
	"github.com/pyropy/registry/ipfs"
	"github.com/rs/zerolog/log"
	"io"
	"math/big"
)

type Registry struct {
	cfg              *config.Config
	ipfsClient       *ipfs.Client
	ethClient        *eth.Client
	registryContract *contract.FileRegistry
}

var ZeroGwei = big.NewFloat(0.0)

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

	r := Registry{
		cfg:              cfg,
		ipfsClient:       ipfsClient,
		ethClient:        client,
		registryContract: c,
	}

	return &r, nil
}

func (r *Registry) UploadFile(ctx context.Context, filePath string, file io.Reader) (string, error) {
	cid, err := r.ipfsClient.UploadFile(ctx, file)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to ipfs: %s", err)
	}

	txOpts, err := r.ethClient.NewTransactOpts(ctx, r.cfg.EthGasLimit, nil, ZeroGwei)
	if err != nil {
		return "", fmt.Errorf("failed to create tx options: %s", err)
	}

	tx, err := r.registryContract.Save(txOpts, filePath, cid)
	if err != nil {
		return "", fmt.Errorf("failed to save cid: %s", err)
	}

	log.Ctx(ctx).Info().
		Str("txHash", tx.Hash().Hex()).
		Msg("transaction submitted")

	return cid, nil
}

func (r *Registry) GetFileCid(ctx context.Context, filePath string) (string, error) {
	callOpts, err := r.ethClient.NewCallOpts(ctx)
	if err != nil {
		return "", err
	}

	return r.registryContract.Get(callOpts, filePath)
}
