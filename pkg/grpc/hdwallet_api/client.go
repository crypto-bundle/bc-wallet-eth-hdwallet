/*
 * MIT License
 *
 * Copyright (c) 2021-2023 Aleksei Kotelnikov
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package hdwallet_api

import (
	"context"
	"errors"

	pbApi "github.com/crypto-bundle/bc-wallet-eth-hdwallet/pkg/grpc/hdwallet_api/proto"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	originGRPC "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	ErrUnableDecodeGrpcErrorStatus = errors.New("unable to decode grpc error status")
)

type Client struct {
	cfg clientConfig

	grpcConn *originGRPC.ClientConn
	client   pbApi.HdWalletApiClient
}

// Init bcexplorer service
// nolint:revive // fixme (autofix)
func (s *Client) Init(ctx context.Context) error {
	options := DefaultDialOptions()
	msgSizeOptions := originGRPC.WithDefaultCallOptions(
		originGRPC.MaxCallRecvMsgSize(DefaultClientMaxReceiveMessageSize),
		originGRPC.MaxCallSendMsgSize(DefaultClientMaxSendMessageSize),
	)
	options = append(options, msgSizeOptions,
		originGRPC.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)

	grpcConn, err := originGRPC.Dial(s.cfg.GetHdWalletServerAddress(), options...)
	if err != nil {
		return err
	}
	s.grpcConn = grpcConn

	s.client = pbApi.NewHdWalletApiClient(grpcConn)

	return nil
}

// Shutdown bcexplorer service
// nolint:revive // fixme (autofix)
func (s *Client) Shutdown(ctx context.Context) error {
	return s.grpcConn.Close()
}

// GetEnabledWallets is function for getting address from bcexplorer
func (s *Client) GetEnabledWallets(ctx context.Context) (*pbApi.GetEnabledWalletsResponse, error) {
	request := &pbApi.GetEnabledWalletsRequest{}

	enabledWallets, err := s.client.GetEnabledWallets(ctx, request)
	if err != nil {
		return nil, err
	}

	return enabledWallets, nil
}

// GetDerivationAddress is function for getting address from bcexplorer
func (s *Client) GetDerivationAddress(ctx context.Context,
	walletUUID string,
	accountIndex uint32,
	internalIndex uint32,
	addressIndex uint32,
) (*pbApi.DerivationAddressResponse, error) {
	request := &pbApi.DerivationAddressRequest{
		AddressIdentity: &pbApi.DerivationAddressIdentity{
			AccountIndex:  accountIndex,
			InternalIndex: internalIndex,
			AddressIndex:  addressIndex,
		},
	}

	md := metadata.New(map[string]string{
		"wallet_uuid": walletUUID,
	})

	requestCtx := metadata.NewOutgoingContext(ctx, md)

	address, err := s.client.GetDerivationAddress(requestCtx, request)
	if err != nil {
		grpcStatus, ok := status.FromError(err)
		if !ok {
			return nil, ErrUnableDecodeGrpcErrorStatus
		}

		switch grpcStatus.Code() {
		case codes.NotFound:
			return nil, nil
		default:
			return nil, err
		}
	}

	return address, nil
}

// nolint:revive // fixme
func NewClient(ctx context.Context,
	cfg clientConfig,
) (*Client, error) {
	srv := &Client{
		cfg: cfg,
	}

	return srv, nil
}
