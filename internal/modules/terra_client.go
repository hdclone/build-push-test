package modules

import (
	"context"
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"broadcaster/internal/config"
)

type tokenAuth struct {
	m map[string]string
}

func (t tokenAuth) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return t.m, nil
}

func (t tokenAuth) RequireTransportSecurity() bool {
	return true
}

func TerraConn(chainConfig *config.ChainConfig) *grpc.ClientConn {
	return Register(fmt.Sprintf("terra_client_%d", chainConfig.ID), func(s string) (Module, error) {
		endpoint := chainConfig.Endpoints[0]
		if len(endpoint) == 0 {
			return nil, fmt.Errorf("undefined chain endpoint, chain: %d", chainConfig.ID)
		}
		authToken := chainConfig.Terra.Auth
		if len(authToken) == 0 {
			return nil, fmt.Errorf("undefined terra grpc token")
		}
		auth := tokenAuth{m: map[string]string{"authorization": authToken}}
		grpcConn, err := grpc.DialContext(
			context.Background(),
			endpoint,
			grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})),
			grpc.WithPerRPCCredentials(auth),
		)
		if err != nil {
			return nil, err
		}

		return grpcConn, nil
	}).(*grpc.ClientConn)
}
