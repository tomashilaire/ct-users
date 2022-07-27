package authgrpcintrcp

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"root/pkg/security"
)

type AuthInterceptor struct {
	jwtManager        security.Security
	accessibleMethods map[string][]string
}

func NewAuthInterceptor(jwtManager security.Security, accessibleMethods map[string][]string) *AuthInterceptor {
	return &AuthInterceptor{jwtManager, accessibleMethods}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) (string, error) {
	_, ok := interceptor.accessibleMethods[method]
	if !ok {
		// everyone can access
		return "", nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := interceptor.jwtManager.NewTokenPayload(accessToken)
	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return claims.UserId, nil
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)

		authId, err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			md.Append("authId", authId)
		}
		newCtx := metadata.NewIncomingContext(ctx, md)
		return handler(newCtx, req)
	}
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println("--> stream interceptor: ", info.FullMethod)

		_, err := interceptor.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		return handler(srv, stream)
	}
}
