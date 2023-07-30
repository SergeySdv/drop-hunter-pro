package zksynclite

import (
	"context"

	//"github.com/zksync-sdk/zksync-sdk-go"
	"github.com/ethereum/go-ethereum/common"
)

type ActivateAccount struct {
	Pk string
}

type activateAccountReq struct {
	Type      string         `json:"type"` // ChangePubKey
	AccountId int            `json:"accountId"`
	Account   common.Address `json:"account"`

	ValidFrom  int `json:"validFrom"`  // 0
	ValidUntil int `json:"validUntil"` // 4294967295
	Nonce      int `json:"nonce"`
	FeeToken   int `json:"feeToken"` // 0

	EthAuthData  *string `json:"ethAuthData"`
	Fee          string  `json:"fee"`
	NewPkHash    string  `json:"newPkHash"`
	EthSignature string  `json:"ethSignature"`
	Signature    Signature
}

type Signature struct {
	PubKey    string `json:"pubKey"`
	Signature string `json:"signature"`
}

//signature?: Signature;

func (c *Client) ActivateAccount(ctx context.Context, req *ActivateAccount) error {
	//txr, err := defi.NewWalletTransactor(req.Pk)
	//if err != nil {
	//	return err
	//}
	//
	//acc, err := c.AccountInfo(ctx, &AccountInfoReq{Addr: txr.WalletAddr})
	//if err != nil {
	//	return err
	//}

	//seed := make([]byte, 32)
	//privateKey, err := zkscrypto.NewPrivateKey(seed)
	//if err != nil {
	//	return err
	//}
	//publicKey, err := privateKey.PublicKey()
	//if err != nil {
	//	return err
	//}
	//publicKeyHash, err := publicKey.Hash()
	//if err != nil {
	//	return err
	//}

	//ethtrx, err := defi.NewWalletTransactor(req.Pk)
	//if err != nil {
	//	return err
	//}

	//ethSignature, err := ethtrx.PrivateKey.Sign([]byte{})
	//if err != nil {
	//	return err
	//}

	//signature, err := privateKey.Sign([]byte{})
	//if err != nil {
	//	return err
	//}

	//data := &activateAccountReq{
	//	Type:        "ChangePubKey",
	//	AccountId:   acc.Id,
	//	Account:     txr.WalletAddr,
	//	ValidFrom:   0,
	//	ValidUntil:  4294967295,
	//	FeeToken:    0,
	//	Nonce:       acc.Committed.Nonce,
	//	EthAuthData: nil,
	//	Fee:         "0", //???
	//	NewPkHash:   publicKeyHash.HexString(),
	//	//EthSignature: string(ethSignature),
	//	Signature: Signature{
	//		PubKey:    signature.HexString(),
	//		Signature: "",
	//	},
	//}
	//println(nil == data)

	return nil
}
