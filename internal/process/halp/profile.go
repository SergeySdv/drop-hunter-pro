package halp

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/zksyncera"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/settings"
	"github.com/hardstylez72/cry/internal/uniclient"
)

type Halp struct {
	profileRepository repository.ProfileRepository
	settingsService   *settings.Service
}

func NewHalp(profileRepository repository.ProfileRepository,
	settingsService *settings.Service) *Halp {
	return &Halp{
		profileRepository: profileRepository,
		settingsService:   settingsService,
	}
}

type Profile struct {
	UserAgent   string
	UserId      string
	Id          string
	Settings    *v1.Settings
	DB          *repository.Profile
	ProxyString string
	WalletAddr  common.Address
	WalletPK    string
	Wallet      *defi.WalletTransactor
}

func (h *Halp) Profile(ctx context.Context, profileId string) (*Profile, error) {

	profile, err := h.profileRepository.GetProfile(ctx, profileId)
	if err != nil {
		return nil, err
	}

	proxyString := ""
	if profile.Proxy.Valid {
		proxyString = profile.Proxy.String
	}

	stgs, err := h.settingsService.GetSettings(ctx, profile.UserId)
	if err != nil {
		return nil, err
	}

	tx, err := defi.NewWalletTransactor(string(profile.MmskPk))
	if err != nil {
		return nil, err
	}
	return &Profile{
		UserAgent:   profile.UserAgent,
		Settings:    stgs,
		DB:          profile,
		ProxyString: proxyString,
		WalletAddr:  tx.WalletAddr,
		WalletPK:    tx.PK,
		Wallet:      tx,
		Id:          profileId,
		UserId:      profile.UserId,
	}, nil
}

func (p *Profile) BaseConfig(network v1.Network) *uniclient.BaseClientConfig {
	n, _ := settings.GetSettingsNetworkSource(network, p.Settings)
	return &uniclient.BaseClientConfig{
		ProxyString:     p.ProxyString,
		RPCEndpoint:     n.RpcEndpoint,
		UserAgentHeader: p.DB.UserAgent,
	}
}

func (p *Profile) EraWallet(network v1.Network) (*zksyncera.WalletTransactor, error) {

	n, err := settings.GetSettingsNetworkSource(network, p.Settings)
	if err != nil {
		return nil, err
	}

	wallet, err := zksyncera.NewWalletTransactor(string(p.DB.MmskPk), big.NewInt(n.Id))
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (p *Profile) EthWallet() (*defi.WalletTransactor, error) {

	wallet, err := defi.NewWalletTransactor(string(p.DB.MmskPk))
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (p *Profile) GetWalletAddr(n v1.Network) (*common.Address, error) {
	var walletAddr common.Address
	if n == v1.Network_ZKSYNCERA || n == v1.Network_ZKSYNCLITE {
		w, err := p.EraWallet(n)
		if err != nil {
			return nil, err
		}
		walletAddr = w.WalletAddr
	} else {
		w, err := p.EthWallet()
		if err != nil {
			return nil, err
		}
		walletAddr = w.WalletAddr
	}
	return &walletAddr, nil
}
