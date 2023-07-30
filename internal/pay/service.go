package pay

import (
	client "github.com/hardstylez72/cry-pay/proto/gen/go/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	cfg *Config
	client.FundsServiceClient
	client.OrdersServiceClient
}

type Config struct {
	Host string
}

func NewService(c *Config) (*Service, error) {
	conn, err := grpc.Dial(c.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	
	return &Service{
		cfg:                 c,
		FundsServiceClient:  client.NewFundsServiceClient(conn),
		OrdersServiceClient: client.NewOrdersServiceClient(conn),
	}, err
}
