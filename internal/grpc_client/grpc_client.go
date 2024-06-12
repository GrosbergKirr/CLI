package grpc_client

import (
	"context"
	"fmt"
	servV1 "github.com/GrosbergKirr/proto_contracts/gen/go/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Api servV1.GatewayServiceClient
}

func NewClient(addr string) (*Client, error) {

	cconn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("failed to create grpc grpc_client: %w", err)
	}

	client := servV1.NewGatewayServiceClient(cconn)

	return &Client{Api: client}, nil

}

func (c *Client) ChangeHostName(ctx context.Context, newname string, addr string, pass string) (string, error) {
	resp, err := c.Api.ChangeHostName(ctx, &servV1.HostRequest{
		NewHostName: newname, Addr: addr, Password: pass,
	})
	if err != nil {
		return "error", err
	}
	return resp.Result, nil
}

func (c *Client) DNSChange(ctx context.Context, dnsname string, addr string, pass string) (string, error) {
	resp, err := c.Api.DNSChange(ctx, &servV1.DNSRequest{
		NewDNSName: dnsname, Addr: addr, Password: pass,
	})
	if err != nil {
		return "error", err
	}
	return resp.Result, nil
}
