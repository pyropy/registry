package testing

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
)

// Backend represents behavior for interacting with an ethereum node.
type Backend interface {
	bind.ContractBackend
	bind.DeployBackend
	Network() string
	ChainID() *big.Int
}

// SimulatedBackend represents a simulated connection to an ethereum node.
type SimulatedBackend struct {
	simulated.Client
	*simulated.Backend
	AutoCommit  bool
	PrivateKeys []*ecdsa.PrivateKey
	chainID     *big.Int
}

// NewSimulatedBackend constructs a simulated backend and a set of private keys
// registered to the backend with a balance of 100 ETH. Use these private keys
// with the NewSimulation call to get an Ethereum API value.
func NewSimulatedBackend(numAccounts int, autoCommit bool, accountBalance *big.Int) (*SimulatedBackend, error) {
	keys := make([]*ecdsa.PrivateKey, numAccounts)
	alloc := make(types.GenesisAlloc)

	for i := 0; i < numAccounts; i++ {
		privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
		if err != nil {
			return nil, fmt.Errorf("unable to generate private key: %w", err)
		}

		keys[i] = privateKey

		alloc[crypto.PubkeyToAddress(privateKey.PublicKey)] = types.Account{
			Balance: big.NewInt(0).Mul(accountBalance, big.NewInt(1e18)),
		}
	}

	gasLimit := uint64(9223372036854775807)
	backend := simulated.NewBackend(alloc, simulated.WithBlockGasLimit(gasLimit))

	// Setting the clock 5.15 minutes into the past to deal with bugs using
	// the simulated clock. Transactions will not be committed otherwise.
	now := time.Since(time.Date(1970, time.January, 1, 0, 5, 15, 0, time.UTC))
	backend.AdjustTime(now)

	backend.Commit()

	b := SimulatedBackend{
		Client:      backend.Client(),
		Backend:     backend,
		AutoCommit:  autoCommit,
		PrivateKeys: keys,
		chainID:     big.NewInt(1337),
	}

	return &b, nil
}

func (b *SimulatedBackend) Network() string {
	return "simulated"
}

// ChainID returns the chain id the backend is connected to.
func (b *SimulatedBackend) ChainID() *big.Int {
	return b.chainID
}

// SendTransaction functions pipes its parameters to the embedded backend and
// also calls Commit() if sb.AutoCommit==true.
func (b *SimulatedBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	if err := b.Client.SendTransaction(ctx, tx); err != nil {
		return err
	}

	if b.AutoCommit {
		b.Commit()
	}

	return nil
}

// SetTime creates a time shift to the simulated clock. It can only be called
// on empty blocks.
func (b *SimulatedBackend) SetTime(t time.Time) {
	b.AdjustTime(time.Since(t))
	b.Commit()
}
