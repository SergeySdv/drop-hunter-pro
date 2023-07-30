package repository

import (
	"context"
	"crypto/sha256"

	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

type ProfileRepositoryCrypto struct {
	source         ProfileRepository
	userRepository UserRepository
	lazanya        string
}

func NewProfileRepositoryCrypto(source ProfileRepository, userRepository UserRepository, lazanya string) ProfileRepository {
	return &ProfileRepositoryCrypto{
		source:         source,
		userRepository: userRepository,
		lazanya:        lazanya,
	}
}

func (c *ProfileRepositoryCrypto) CreateProfile(ctx context.Context, req *Profile) error {

	mmskPk, err := lib.Encrypt(req.UserId, c.lazanya, req.MmskPk)
	if err != nil {
		return err
	}

	h := sha256.New()
	h.Write(req.MmskId)
	req.MmskId = h.Sum(nil)

	req.MmskPk = mmskPk

	return c.source.CreateProfile(ctx, req)
}

func (c *ProfileRepositoryCrypto) SearchNotConnectedOkexDepositProfile(ctx context.Context, userId string) ([]Profile, error) {
	list, err := c.source.SearchNotConnectedOkexDepositProfile(ctx, userId)
	if err != nil {
		return nil, err
	}

	for i := range list {
		list[i].MmskPk, err = lib.Decrypt(userId, c.lazanya, list[i].MmskPk)
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}

func (c *ProfileRepositoryCrypto) ListProfiles(ctx context.Context, userId string) ([]Profile, error) {

	list, err := c.source.ListProfiles(ctx, userId)
	if err != nil {
		return nil, err
	}

	for i := range list {
		list[i].MmskPk, err = lib.Decrypt(userId, c.lazanya, list[i].MmskPk)
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}

func (c *ProfileRepositoryCrypto) ExportProfiles(ctx context.Context, userId string) ([]Profile, error) {

	list, err := c.source.ExportProfiles(ctx, userId)
	if err != nil {
		return nil, err
	}

	for i := range list {
		list[i].MmskPk, err = lib.Decrypt(userId, c.lazanya, list[i].MmskPk)
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}

func (c *ProfileRepositoryCrypto) DeleteProfile(ctx context.Context, req *v1.DeleteProfileRequest) (*v1.DeleteProfileResponse, error) {
	return c.source.DeleteProfile(ctx, req)
}

func (c *ProfileRepositoryCrypto) SearchProfile(ctx context.Context, pattern, userId string) ([]Profile, error) {
	list, err := c.source.SearchProfile(ctx, pattern, userId)
	if err != nil {
		return nil, err
	}
	for i := range list {
		list[i].MmskPk, err = lib.Decrypt(userId, c.lazanya, list[i].MmskPk)
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}

func (c *ProfileRepositoryCrypto) GetProfile(ctx context.Context, id string) (*Profile, error) {
	p, err := c.source.GetProfile(ctx, id)
	if err != nil {
		return nil, err
	}

	p.MmskPk, err = lib.Decrypt(p.UserId, c.lazanya, p.MmskPk)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (c *ProfileRepositoryCrypto) ValidateLabel(ctx context.Context, request *ValidateLabelReq) (*bool, error) {
	return c.source.ValidateLabel(ctx, request)
}

func (c *ProfileRepositoryCrypto) UpdateProfile(ctx context.Context, req *Profile) error {
	return c.source.UpdateProfile(ctx, req)
}
