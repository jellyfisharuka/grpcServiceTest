package auth

import (
	"context"
	ssov1 "protos/gen/go/sso"

	"google.golang.org/grpc"
)

type ServerAPI struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &ServerAPI{}) //регистрирует обработчик
}
func (s *ServerAPI) Login(ctx context.Context, req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	return &ssov1.LoginResponse{
		Token: req.GetEmail(),//можно и просто req.Email() разницы нет, но чисто для интереса вернуть геттеры 
	}, nil
}
func (s *ServerAPI) Register(ctx context.Context, req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	panic("implement me")
}
func (s *ServerAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	panic("implement me")
}
