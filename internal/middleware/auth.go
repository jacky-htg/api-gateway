package middleware

import (
	"context"
	users "inventory-gateway-service/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Auth struct
type Auth struct {
	Client users.AuthServiceClient
	Token  string
}

// Unary interceptor
func (u *Auth) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		var err error

		if info.FullMethod != "/wiradata.users.AuthService/Login" &&
			info.FullMethod != "/wiradata.users.AuthService/ForgotPassword" &&
			info.FullMethod != "/wiradata.users.AuthService/ResetPassword" &&
			info.FullMethod != "/wiradata.users.CompanyService/Registration" {
			err = u.auth(ctx, info.FullMethod)
			if err != nil {
				return nil, err
			}
		}

		return handler(ctx, req)
	}
}

// Stream interceptor
func (u *Auth) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println("--> stream interceptor: ", info.FullMethod)
		err := u.auth(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		return handler(srv, stream)
	}
}

func (u *Auth) auth(ctx context.Context, fullMethod string) error {
	isAuth, err := u.Client.IsAuth(ctx, &users.String{String_: fullMethod})
	if err != nil {
		return err
	}

	if !isAuth.Boolean {
		return status.Error(codes.Unauthenticated, "You dont have permission")
	}

	return nil
}