package config

import (
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
			want: DefaultConfig,
		},
		{
			name: "custom",
			env: map[string]string{
				EnvPort:        "8000",
				EnvIpfsPort:    "8081",
				EnvIpfsHost:    "localhost",
				EnvIpfsAutoPin: "false",
			},
			want: &Config{
				Port:        8000,
				IpfsPort:    "8081",
				IpfsHost:    "localhost",
				EthRpcUrl:   "http://localhost:8545",
				IpfsAutoPin: false,
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
