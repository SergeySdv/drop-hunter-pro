package repository

import (
	"context"

	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type WithdrawerRepositoryCrypto struct {
	source         WithdrawerRepository
	userRepository UserRepository
	lazanya        string
}

func NewWithdrawerRepositoryCrypto(source WithdrawerRepository, userRepository UserRepository, lazanya string) WithdrawerRepository {
	return &WithdrawerRepositoryCrypto{
		source:         source,
		userRepository: userRepository,
		lazanya:        lazanya,
	}
}

func (c *WithdrawerRepositoryCrypto) GetAttachedAddr(ctx context.Context, profileId, userId string) (*OkexDepositAddr, error) {
	return c.source.GetAttachedAddr(ctx, profileId, userId)
}

func (c *WithdrawerRepositoryCrypto) ListDepositAddrAttach(ctx context.Context, userId string) (map[Addr]OkexDepositAddr, error) {
	return c.source.ListDepositAddrAttach(ctx, userId)
}

func (c *WithdrawerRepositoryCrypto) OkexDepositAddrAttach(ctx context.Context, req *OkexDepositAddr) error {
	return c.source.OkexDepositAddrAttach(ctx, req)
}

func (c *WithdrawerRepositoryCrypto) OkexDepositAddrDetach(ctx context.Context, req *OkexDepositAddr) error {
	return c.source.OkexDepositAddrDetach(ctx, req)
}

func (c *WithdrawerRepositoryCrypto) CreateWithdrawer(ctx context.Context, req *Withdrawer) (*Withdrawer, error) {
	SecretKey, err := lib.Encrypt(req.UserId, c.lazanya, req.SecretKey)
	if err != nil {
		return nil, err
	}

	ApiKey, err := lib.Encrypt(req.UserId, c.lazanya, req.ApiKey)
	if err != nil {
		return nil, err
	}

	req.SecretKey = SecretKey
	req.ApiKey = ApiKey

	return c.source.CreateWithdrawer(ctx, req)
}

func (c *WithdrawerRepositoryCrypto) CreateSubWithdrawer(ctx context.Context, req *Withdrawer) error {
	SecretKey, err := lib.Encrypt(req.UserId, c.lazanya, req.SecretKey)
	if err != nil {
		return err
	}

	ApiKey, err := lib.Encrypt(req.UserId, c.lazanya, req.ApiKey)
	if err != nil {
		return err
	}

	req.SecretKey = SecretKey
	req.ApiKey = ApiKey

	return c.source.CreateSubWithdrawer(ctx, req)
}

func (c *WithdrawerRepositoryCrypto) ListWithdrawers(ctx context.Context, userId string) ([]Withdrawer, error) {
	return c.source.ListWithdrawers(ctx, userId)
}

func (c *WithdrawerRepositoryCrypto) ExportExchangeAccounts(ctx context.Context, userId string) ([]Withdrawer, error) {
	res, err := c.source.ExportExchangeAccounts(ctx, userId)
	if err != nil {
		return nil, err
	}

	for i := range res {
		w := &res[i]

		apiKey, err := lib.Decrypt(w.UserId, c.lazanya, w.ApiKey)
		if err != nil {
			return nil, err
		}
		w.ApiKey = apiKey
		secretKey, err := lib.Decrypt(w.UserId, c.lazanya, w.SecretKey)
		if err != nil {
			return nil, err
		}
		w.SecretKey = secretKey

	}

	return res, nil
}

func (c *WithdrawerRepositoryCrypto) ListSubWithdrawers(ctx context.Context, id, userId string) ([]Withdrawer, error) {
	return c.source.ListSubWithdrawers(ctx, id, userId)
}

func (c *WithdrawerRepositoryCrypto) DeleteWithdrawer(ctx context.Context, req *v1.DeleteWithdrawerRequest) (*v1.DeleteWithdrawerResponse, error) {
	return c.source.DeleteWithdrawer(ctx, req)
}

func (c *WithdrawerRepositoryCrypto) UpdateWithdrawer(ctx context.Context, req *Withdrawer) error {
	return c.source.UpdateWithdrawer(ctx, req)
}

func (c *WithdrawerRepositoryCrypto) GetWithdrawer(ctx context.Context, id, userId string) (*Withdrawer, error) {
	p, err := c.source.GetWithdrawer(ctx, id, userId)
	if err != nil {
		return nil, err
	}

	apiKey, err := lib.Decrypt(p.UserId, c.lazanya, p.ApiKey)
	if err != nil {
		return nil, err
	}
	p.ApiKey = apiKey
	secretKey, err := lib.Decrypt(p.UserId, c.lazanya, p.SecretKey)
	if err != nil {
		return nil, err
	}
	p.SecretKey = secretKey

	return p, nil
}
