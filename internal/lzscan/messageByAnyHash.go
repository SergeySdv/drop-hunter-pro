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
	queryGetMessageByAnyHash = `{
   "query":"query getMessageByAnyHash($hash: String!) {\n  dstTx: messagesByDstTxHash(dstTxHash: $hash) {\n    items {\n      ...MessageDetails\n      __typename\n    }\n    __typename\n  }\n  srcTx: messagesBySrcTxHash(srcTxHash: $hash) {\n    items {\n      ...MessageDetails\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment MessageDetails on Message {\n  srcUaAddress\n  dstUaAddress\n  updated\n  created\n  srcChainId\n  dstChainId\n  dstTxHash\n  dstTxError\n  srcTxHash\n  srcBlockHash\n  srcBlockNumber\n  srcUaNonce\n  status\n  adapterParams\n}",
   "operationName":"getMessageByAnyHash",
   "variables":{
      "hash":"%s"
   }
}`
)

func (s *Service) GetMessageByAnyHash(ctx context.Context, hash string) (*GetMessageByAnyHashResponse, error) {

	body := bytes.NewBuffer([]byte(fmt.Sprintf(queryGetMessageByAnyHash, hash)))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.Host, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := s.HttpCli.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		b, _ := io.ReadAll(res.Body)
		return nil, errors.New(string(b))
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response GetMessageByAnyHashResponse
	if err := json.Unmarshal(b, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type GetMessageByAnyHashResponse struct {
	Data struct {
		DstTx struct {
			Items    []interface{} `json:"items"`
			Typename string        `json:"__typename"`
		} `json:"dstTx"`
		SrcTx struct {
			Items []struct {
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
			} `json:"items"`
			Typename string `json:"__typename"`
		} `json:"srcTx"`
	} `json:"data"`
}
