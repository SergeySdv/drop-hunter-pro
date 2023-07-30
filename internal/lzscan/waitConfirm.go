package lzscan

import (
	"context"
	"encoding/json"
	"time"

	"github.com/hardstylez72/cry/internal/log"
	"github.com/pkg/errors"
)

const (
	StatusINFLIGHT  = "INFLIGHT"
	StatusDELIVERED = "DELIVERED"
	StatusFAILED    = "FAILED"
)

func (s *Service) WaitConfirm(ctx context.Context, txId string) (string, error) {

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return StatusINFLIGHT, nil
		case <-ticker.C:
			tx, err := s.GetTxById(ctx, txId)
			if err != nil {
				log.Log.Warn("lzscan.WaitConfirm: " + err.Error())
				continue
			}

			switch tx.Data.Message.Status {
			case StatusINFLIGHT:
				continue
			case StatusFAILED:
				b, _ := json.Marshal(&tx.Data.Message.DstTxError)
				return StatusFAILED, errors.New(string(b))
			case StatusDELIVERED:
				return StatusDELIVERED, nil
			}
		}
	}
}

func (s *Service) GetTxById(ctx context.Context, txId string) (*GetMessageChainIdsSrcUAAddressAndSrcOutboundNonceResponse, error) {
	message, err := s.GetMessageByAnyHash(ctx, txId)
	if err != nil {
		return nil, errors.Wrap(err, "GetMessageByAnyHash")
	}

	if len(message.Data.SrcTx.Items) == 0 {
		return nil, ErrNotFound
	}

	src := message.Data.SrcTx.Items[0]
	res, err := s.GetMessageChainIdsSrcUAAddressAndSrcOutboundNonce(ctx, &GetMessageChainIdsSrcUAAddressAndSrcOutboundNonceRequest{
		SrcChainId:   src.SrcChainId,
		DstChainId:   src.DstChainId,
		SrcUaAddress: src.SrcUaAddress,
		DstUaAddress: src.DstUaAddress,
		SrcUaNonce:   src.SrcUaNonce,
	})

	if err != nil {
		return nil, errors.Wrap(err, "GetMessageChainIdsSrcUAAddressAndSrcOutboundNonce")
	}

	return res, nil
}
