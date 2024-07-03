package config

import (
	"github.com/ethereum/go-ethereum/common"
	"os"
	"reflect"
	"testing"
)

func TestLoadConfigFromEnv(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
		want *Config
	}{
		{
			name: "default",
			env:  map[string]string{},
			want: &Config{
				Port:                        8000,
				IpfsUrl:                     "http://localhost:5001",
				IpfsAutoPin:                 true,
				EthRpcUrl:                   "http://localhost:8545",
				EthPrivateKey:               nil,
				EthGasLimit:                 21000,
				FileRegistryContractAddress: common.HexToAddress("0x"),
			},
		},
		{
			name: "custom",
			env: map[string]string{
				"HTTP_PORT":     "8080",
				"IPFS_URL":      "http://localhost:8081",
				"IPFS_AUTO_PIN": "false",
			},
			want: &Config{
				Port:                        8080,
				IpfsUrl:                     "http://localhost:8081",
				IpfsAutoPin:                 false,
				EthRpcUrl:                   "http://localhost:8545",
				EthPrivateKey:               nil,
				EthGasLimit:                 21000,
				FileRegistryContractAddress: common.HexToAddress("0x"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set env variables
			for k, v := range tt.env {
				if err := os.Setenv(k, v); err != nil {
					t.Fatalf("cannot set env var: %s", err)
				}
			}

			got, err := LoadConfigFromEnv()
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
