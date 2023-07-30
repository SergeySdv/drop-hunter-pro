package lzscan

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const (
	queryGetMessageChainIdsSrcUAAddressAndSrcOutboundNonce = `
	{
   "query":"query getMessageChainIdsSrcUAAddressAndSrcOutboundNonce($srcChainId: Int!, $dstChainId: Int!, $srcUaAddress: String!, $srcUaNonce: Int!, $dstUaAddress: String) {\n  message(\n    srcChainId: $srcChainId\n    dstChainId: $dstChainId\n    srcUaAddress: $srcUaAddress\n    srcUaNonce: $srcUaNonce\n    dstUaAddress: $dstUaAddress\n  ) {\n    ...MessageDetails\n    __typename\n  }\n}\n\nfragment MessageDetails on Message {\n  srcUaAddress\n  dstUaAddress\n  updated\n  created\n  srcChainId\n  dstChainId\n  dstTxHash\n  dstTxError\n  srcTxHash\n  srcBlockHash\n  srcBlockNumber\n  srcUaNonce\n  status\n  adapterParams\n}",
   "operationName":"getMessageChainIdsSrcUAAddressAndSrcOutboundNonce",
   "variables":{
      "srcChainId":%d,
      "dstChainId":%d,
      "srcUaAddress":"%s",
      "dstUaAddress":"%s",
      "srcUaNonce":%d
   }
}
`
)

type GetMessageChainIdsSrcUAAddressAndSrcOutboundNonceRequest struct {
	SrcChainId   int
	DstChainId   int
	SrcUaAddress string
	DstUaAddress string
	SrcUaNonce   int
}

type GetMessageChainIdsSrcUAAddressAndSrcOutboundNonceResponse struct {
	Data struct {
		Message struct {
			SrcUaAddress   string      `json:"srcUaAddress"`
			DstUaAddress   string      `json:"dstUaAddress"`
			Updated        int         `json:"updated"`
			Created        int         `json:"created"`
			SrcChainId     int         `json:"srcChainId"`
			DstChainId     int         `json:"dstChainId"`
			DstTxHash      string      `json:"dstTxHash"`
			DstTxError     interface{} `json:"dstTxError"`
			SrcTxHash      string      `json:"srcTxHash"`
			SrcBlockHash   string      `json:"srcBlockHash"`
			SrcBlockNumber string      `json:"srcBlockNumber"`
			SrcUaNonce     int         `json:"srcUaNonce"`
			Status         string      `json:"status"`
			AdapterParams  string      `json:"adapterParams"`
			Typename       string      `json:"__typename"`
		} `json:"message"`
	} `json:"data"`
}

func (s *Service) GetMessageChainIdsSrcUAAddressAndSrcOutboundNonce(ctx context.Context, req *GetMessageChainIdsSrcUAAddressAndSrcOutboundNonceRequest) (*GetMessageChainIdsSrcUAAddressAndSrcOutboundNonceResponse, error) {

	body := bytes.NewBuffer([]byte(fmt.Sprintf(queryGetMessageChainIdsSrcUAAddressAndSrcOutboundNonce,
		req.SrcChainId,
		req.DstChainId,
		req.SrcUaAddress,
		req.DstUaAddress,
		req.SrcUaNonce,
	)))

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, s.Host, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := s.HttpCli.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		b, _ := io.ReadAll(response.Body)
		return nil, errors.New(string(b))
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var res GetMessageChainIdsSrcUAAddressAndSrcOutboundNonceResponse
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
