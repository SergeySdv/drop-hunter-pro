package lzscan

import (
	"net/http"
)

type Service struct {
	HttpCli *http.Client
	Host    string
}

func NewService() *Service {
	return &Service{
		HttpCli: &http.Client{},
		Host:    "https://qtrhqncfi4.execute-api.us-east-1.amazonaws.com/prod/graphql",
	}
}
