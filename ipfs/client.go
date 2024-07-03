package ipfs

import (
	"context"
	"github.com/ipfs/boxo/files"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/ipfs/kubo/core/coreiface/options"
	"github.com/pyropy/registry/config"
	"io"
	"net/http"
)

type Client interface {
	UploadFile(ctx context.Context, file io.Reader) (string, error)
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
	f := files.NewReaderFile(file)
	path, err := c.node.Unixfs().Add(ctx, f, getIPFSUploadSettings(c.cfg))
	if err != nil {
		return "", err
	}

	return path.RootCid().String(), nil
}

func getIPFSUploadSettings(cfg *config.Config) options.UnixfsAddOption {
	return func(settings *options.UnixfsAddSettings) error {
		settings.Pin = cfg.IpfsAutoPin
		return nil
	}
}

type MockClient struct {
}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (c *MockClient) UploadFile(_ context.Context, _ io.Reader) (string, error) {
	return "bafybeigdyrzt5sfp7udm7hu76uh7y26nf3efuylqabf3oclgtqy55fbzdi", nil
}
