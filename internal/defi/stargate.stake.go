package defi

import (
	"context"
	"crypto_scripts/internal/defi/contracts/stargate/stake"
	"crypto_scripts/internal/lib"
	v1 "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

const (
	MaxStakeStimeDuration = 94608000
	MixStakeStimeDuration = 2419200
)

type StakeSTGReq struct {
	Amount *big.Int
	Wallet *WalletTransactor
}

type StakeSTGRes struct {
	Tx *types.Transaction
}

func (c *EtheriumClient) StakeSTG(ctx context.Context, req *StakeSTGReq) (*StakeSTGRes, error) {
	return lib.Retry[*StakeSTGReq, *StakeSTGRes](ctx, c.stakeSTG, req, &lib.RetryOpt{
		RetryCount: RetryMax,
	})
}
func (c *EtheriumClient) stakeSTG(ctx context.Context, req *StakeSTGReq) (*StakeSTGRes, error) {

	viewer, err := stake.NewStakerCaller(c.c.Dict.Stargate.STGStakingAddress, c.cli)
	if err != nil {
		return nil, errors.Wrap(err, "stargate.NewStakeStgCaller")
	}

	locked, err := viewer.Locked(nil, req.Wallet.WalletAddr)
	if err != nil {
		return nil, errors.Wrap(err, "viewer.Locked")
	}

	allowed, err := c.TokenAllowed(ctx, &AllowedReq{
		WalletAddr:  req.Wallet.WalletAddr,
		SpenderAddr: c.c.Dict.Stargate.STGStakingAddress,
		Token:       v1.Token_STG,
	})
	if err != nil {
		return nil, errors.Wrap(err, "STGAllowed")
	}

	want := big.NewInt(0).Add(req.Amount, locked.Amount)
	have := allowed.Allowance

	if want.Cmp(have) == 1 {
		_, err := c.TokenApprove(ctx, &ApproveReq{
			Wallet: req.Wallet,
			amount: want,
			Token:  v1.Token_STG,
		})
		if err != nil {
			return nil, errors.Wrap(err, "STGApprove")
		}
	}

	staker, err := stake.NewStakerTransactor(c.c.Dict.Stargate.STGStakingAddress, c.cli)
	if err != nil {
		return nil, errors.Wrap(err, "stargate.NewStakeStgTransactor")
	}

	chainID, err := c.cli.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	opt, err := bind.NewKeyedTransactorWithChainID(req.Wallet.PrivateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}
	opt.Context = ctx

	var tx *types.Transaction
	if locked.Amount.Int64() == 0 {
		t := time.Now().Add(time.Duration(MaxStakeStimeDuration) * time.Second)
		tx, err = staker.CreateLock(opt, req.Amount, big.NewInt(t.Unix()))
		if err != nil {
			return nil, errors.Wrap(err, "staker.CreateLock")
		}
	} else {
		tx, err = staker.IncreaseAmount(opt, req.Amount)
		if err != nil {
			return nil, errors.Wrap(err, "staker.IncreaseAmountAndTime")
		}
	}

	return &StakeSTGRes{Tx: tx}, nil
}
