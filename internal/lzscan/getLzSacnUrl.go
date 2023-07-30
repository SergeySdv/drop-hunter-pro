package lzscan

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

const (
	UrlLzTxMask = `https://layerzeroscan.com/%d/address/%s/message/%d/address/%s/nonce/%d`
)

var ErrNotFound = errors.New("not found")

func (s *Service) GetTxUrl(ctx context.Context, txId string) (string, error) {
	res, err := s.GetMessageByAnyHash(ctx, txId)
	if err != nil {
		return "", err
	}

	if len(res.Data.SrcTx.Items) == 0 {
		return "", ErrNotFound
	}

	src := res.Data.SrcTx.Items[0]

	return fmt.Sprintf(UrlLzTxMask, src.SrcChainId, src.SrcUaAddress, src.DstChainId, src.DstUaAddress, src.SrcUaNonce), nil
}
