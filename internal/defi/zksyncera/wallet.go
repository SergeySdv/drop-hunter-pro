package zksyncera

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go/accounts"
)

type WalletTransactor struct {
	WalletAddr   common.Address
	WalletAddrHR string
	PK           string
	PKb          []byte
	Signer       *accounts.DefaultEthSigner
}

func NewWalletTransactor(privateKey string, networkId *big.Int) (*WalletTransactor, error) {

	pkb, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "hex.DecodeString")
	}
	signer, err := accounts.NewEthSignerFromRawPrivateKey(pkb, networkId.Int64())
	if err != nil {
		return nil, err
	}

	return &WalletTransactor{
		WalletAddr:   signer.GetAddress(),
		WalletAddrHR: signer.GetAddress().String(),
		PK:           privateKey,
		PKb:          pkb,
		Signer:       signer,
	}, nil
}
