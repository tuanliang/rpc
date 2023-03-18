package server

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	ClientHeaderAccessKey = "client-id"
	ClientHeaderSecretKey = "client-secret"
)

func NewClientCredential(ak, sk string) metadata.MD {
	return metadata.MD{
		ClientHeaderAccessKey: []string{ak},
		ClientHeaderSecretKey: []string{sk},
	}
}

func NewAuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return (&GrpcAuther{}).UnaryServerInterceptor
}
func NewStreamServerInterceptor() grpc.StreamServerInterceptor {
	return (&GrpcAuther{}).StreamServerInterceptor
}

type GrpcAuther struct {
}

// request-response interceptor
func (a *GrpcAuther) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 1.读取凭证，凭证放在meta信息[http2 header]
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("ctx is not an grpc incoming context")
	}

	// 从meta data中获取从客户端传递过来的凭证
	clientId, clientSecret := a.getClientGredentialsFromMeta(md)

	// 校验凭证合法性
	if err := a.validateServiceCredential(clientId, clientSecret); err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

// stream rpc interceptor
func (a *GrpcAuther) StreamServerInterceptor(srv interface{}, ss grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	//
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return fmt.Errorf("ctx is not an grpc incoming context")
	}
	// 从meta data中获取从客户端传递过来的凭证
	clientId, clientSecret := a.getClientGredentialsFromMeta(md)

	// 校验凭证合法性
	if err := a.validateServiceCredential(clientId, clientSecret); err != nil {
		return err
	}
	return handler(srv, ss)
}

func (a *GrpcAuther) getClientGredentialsFromMeta(md metadata.MD) (clientId, clientSecret string) {
	cakList := md[ClientHeaderAccessKey]
	if len(cakList) > 0 {
		clientId = cakList[0]
	}
	cskList := md[ClientHeaderSecretKey]
	if len(cskList) > 0 {
		clientSecret = cskList[0]
	}

	return
}
func (a *GrpcAuther) validateServiceCredential(clientId, clientSecret string) error {
	if !(clientId == "admin" && clientSecret == "123456") {
		fmt.Println(clientId, clientSecret)
		return status.Errorf(codes.Unauthenticated, "client_id or client_secret not correct")
	}
	return nil
}
