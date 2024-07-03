package ipfs

import (
	"context"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/ipfs/kubo/core/coreiface/options"
	"github.com/pyropy/registry/config"
	"io"
	"net/http"
)

type Client interface {
	UploadFile(ctx context.Context, file io.Reader) (string, error)
}

type SimulatedClient struct {
}

func NewSimulatedClient() *SimulatedClient {
	return &SimulatedClient{}
}

func (c *SimulatedClient) UploadFile(_ context.Context, _ io.Reader) (string, error) {
	return "bafybeigdyrzt5sfp7udm7hu76uh7y26nf3efuylqabf3oclgtqy55fbzdi", nil
}

// DialedClient is a thin wrapper around IPFS client
type DialedClient struct {
	cfg  *config.Config
	node *rpc.HttpApi
}

// NewDialedClient creates a new IPFS client
// connected to the IPFS node specified in the config
func NewDialedClient(cfg *config.Config) (*DialedClient, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy:             http.ProxyFromEnvironment,
			DisableKeepAlives: true,
		},
	}

	node, err := rpc.NewURLApiWithClient(cfg.IpfsUrl, client)
	return &DialedClient{cfg: cfg, node: node}, err
}

// UploadFile uploads a file to IPFS and returns CID of the uploaded file
func (c *DialedClient) UploadFile(ctx context.Context, file io.Reader) (string, error) {
	block, err := c.node.Block().Put(ctx, file, getIPFSUploadSettings(c.cfg))
	if err != nil {
		return "", err
	}

	return block.Path().RootCid().String(), nil
}

func getIPFSUploadSettings(cfg *config.Config) options.BlockPutOption {
	return func(settings *options.BlockPutSettings) error {
		settings.Pin = cfg.IpfsAutoPin
		return nil
	}
}
