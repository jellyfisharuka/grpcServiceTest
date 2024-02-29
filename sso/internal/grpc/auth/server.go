package auth

import (
	"context"
	ssov1 "protos/gen/go/sso"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
type Server interface {
	
}
type ServerAPI struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &ServerAPI{}) //регистрирует обработчик
}
const (
	emplyValue=0
)
func (s *ServerAPI) Login(ctx context.Context, req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is requried")
	}
	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}
	if req.GetAppId() == emplyValue {
       return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}
  //TODO: implement login via auth service
	return &ssov1.LoginResponse{
		Token: req.GetEmail(), //можно и просто req.Email() разницы нет, но чисто для интереса вернуть геттеры
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
