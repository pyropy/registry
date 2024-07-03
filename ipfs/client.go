package ipfs

import (
	"context"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/ipfs/kubo/core/coreiface/options"
	"github.com/pyropy/registry/config"
	"io"
	"net/http"
)

// Client is a thin wrapper around IPFS client
type Client struct {
	cfg  *config.Config
	node *rpc.HttpApi
}

func NewClient(cfg *config.Config) (*Client, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy:             http.ProxyFromEnvironment,
			DisableKeepAlives: true,
		},
	}

	node, err := rpc.NewURLApiWithClient(cfg.IpfsUrl, client)
	return &Client{cfg: cfg, node: node}, err
}

// UploadFile uploads a file to IPFS and returns CID of the uploaded file
func (c *Client) UploadFile(ctx context.Context, file io.Reader) (string, error) {
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
