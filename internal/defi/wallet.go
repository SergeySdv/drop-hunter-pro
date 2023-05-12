package defi

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

type WalletTransactor struct {
	PrivateKey   *ecdsa.PrivateKey
	PublicKey    *ecdsa.PublicKey
	WalletAddr   common.Address
	WalletAddrHR string
}

func NewWalletTransactor(privateKey string) (*WalletTransactor, error) {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "crypto.HexToECDSA(privateKey)")
	}

	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	walletAddr := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &WalletTransactor{
		PrivateKey:   pk,
		PublicKey:    publicKeyECDSA,
		WalletAddr:   walletAddr,
		WalletAddrHR: walletAddr.String(),
	}, nil
}
