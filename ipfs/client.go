package ipfs

import (
	"context"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/ipfs/kubo/core/coreiface/options"
	"github.com/pyropy/registry/config"
	"io"
	"net"
	"net/http"
)

type Client struct {
	cfg  *config.Config
	node *rpc.HttpApi
}

func NewClient(cfg *config.Config) (*Client, error) {
	url := net.JoinHostPort(cfg.IpfsHost, cfg.IpfsPort)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy:             http.ProxyFromEnvironment,
			DisableKeepAlives: true,
		},
	}

	node, err := rpc.NewURLApiWithClient(url, client)
	return &Client{cfg: cfg, node: node}, err
}

// UploadFile uploads a file to IPFS and returns CID of the uploaded file
func (c *Client) UploadFile(ctx context.Context, file io.Reader) (*path.ImmutablePath, error) {
	block, err := c.node.Block().Put(ctx, file, getIPFSUploadSettings(c.cfg))
	if err != nil {
		return nil, err
	}

	p := block.Path()
	return &p, nil
}

func getIPFSUploadSettings(cfg *config.Config) options.BlockPutOption {
	return func(settings *options.BlockPutSettings) error {
		settings.Pin = cfg.IpfsPin
		return nil
	}
}
