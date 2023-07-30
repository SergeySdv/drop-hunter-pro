package defi

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
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

//func (c *EtheriumClient) StakeSTG(ctx context.Context, req *StakeSTGReq) (*StakeSTGRes, error) {
//	return lib.Retry[*StakeSTGReq, *StakeSTGRes](ctx, c.stakeSTG, req, &lib.RetryOpt{
//		RetryCount: RetryMax,
//	})
//}
//func (c *EtheriumClient) stakeSTG(ctx context.Context, req *StakeSTGReq) (*StakeSTGRes, error) {
//
//	viewer, err := stake.NewStakerCaller(c.Cfg.Dict.Stargate.STGStakingAddress, c.Cli)
//	if err != nil {
//		return nil, errors.Wrap(err, "stargate.NewStakeStgCaller")
//	}
//
//	locked, err := viewer.Locked(nil, req.Wallet.WalletAddr)
//	if err != nil {
//		return nil, errors.Wrap(err, "viewer.Locked")
//	}
//
//	allowed, err := c.TokenAllowed(ctx, &AllowedReq{
//		WalletAddr:  req.Wallet.WalletAddr,
//		SpenderAddr: c.Cfg.Dict.Stargate.STGStakingAddress,
//		Token:       v1.Token_STG,
//	})
//	if err != nil {
//		return nil, errors.Wrap(err, "STGAllowed")
//	}
//
//	want := big.NewInt(0).Add(req.Amount, locked.Amount)
//	have := allowed.Allowance
//
//	if want.Cmp(have) == 1 {
//		_, err := c.TokenApprove(ctx, &ApproveReq{
//			Wallet: req.Wallet,
//			Amount: want,
//			Token:  v1.Token_STG,
//		})
//		if err != nil {
//			return nil, errors.Wrap(err, "STGApprove")
//		}
//	}
//
//	staker, err := stake.NewStakerTransactor(c.Cfg.Dict.Stargate.STGStakingAddress, c.Cli)
//	if err != nil {
//		return nil, errors.Wrap(err, "stargate.NewStakeStgTransactor")
//	}
//
//	chainID, err := c.Cli.ChainID(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	opt, err := bind.NewKeyedTransactorWithChainID(req.Wallet.PrivateKey, chainID)
//	if err != nil {
//		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
//	}
//	opt.Context = ctx
//
//	var Tx *types.Transaction
//	if locked.Amount.Int64() == 0 {
//		t := time.Now().Add(time.Duration(MaxStakeStimeDuration) * time.Second)
//		Tx, err = staker.CreateLock(opt, req.Amount, big.NewInt(t.Unix()))
//		if err != nil {
//			return nil, errors.Wrap(err, "staker.CreateLock")
//		}
//	} else {
//		Tx, err = staker.IncreaseAmount(opt, req.Amount)
//		if err != nil {
//			return nil, errors.Wrap(err, "staker.IncreaseAmountAndTime")
//		}
//	}
//
//	return &StakeSTGRes{Tx: Tx}, nil
//}
