package grpc

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	grpcMetadata "google.golang.org/grpc/metadata"
	"waas-service/auth/jwt"
	"waas-service/auth/jwt/jose"
)

func UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := grpcMetadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("cannot get metadata")
	}

	validator, err := jose.NewValidator()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	JWT := jwt.New(validator)
	if err = JWT.Check(md); err != nil {
		return nil, err
	}
	log.Println("--> unary interceptor: ", info.FullMethod)

	return handler(ctx, req)
}
