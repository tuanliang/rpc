package client

import (
	"context"

	"github.com/tuanliang/rpc/grpc/middleware/server"
)

func NewAuthentication(ak, sk string) *Authentication {
	return &Authentication{
		ClientId:     ak,
		ClientSecret: sk,
	}
}

type Authentication struct {
	ClientId     string
	ClientSecret string
}

func (a *Authentication) build() map[string]string {
	return map[string]string{
		server.ClientHeaderAccessKey: a.ClientId,
		server.ClientHeaderSecretKey: a.ClientSecret,
	}
}

func (a *Authentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return a.build(), nil
}
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
