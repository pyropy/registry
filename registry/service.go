package registry

import (
	"context"
	"github.com/pyropy/registry/config"
	"github.com/pyropy/registry/ipfs"
	"io"
)

type Registry struct {
	ipfsClient *ipfs.Client
}

func NewRegistry(cfg *config.Config) (*Registry, error) {
	ipfsClient, err := ipfs.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &Registry{
		ipfsClient: ipfsClient,
	}, nil
}

func (r *Registry) UploadFile(ctx context.Context, filePath string, file io.Reader) (string, error) {
	cid, err := r.ipfsClient.UploadFile(ctx, file)
	if err != nil {
		return "", err
	}

	// TODO: Upload to eth
	return cid.String(), nil
}

func (r *Registry) GetFileCid(ctx context.Context, filePath string) (string, error) {
	return "", nil
}
