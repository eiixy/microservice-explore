package service

import (
	pb "account/api/account/v1"
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type AccountService struct {
	pb.UnimplementedAccountServer
}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (s *AccountService) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountReply, error) {
	return &pb.CreateAccountReply{}, nil
}
func (s *AccountService) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountReply, error) {
	return &pb.UpdateAccountReply{}, nil
}
func (s *AccountService) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountReply, error) {
	return &pb.DeleteAccountReply{}, nil
}
func (s *AccountService) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountReply, error) {
	return &pb.GetAccountReply{}, nil
}
func (s *AccountService) ListAccount(ctx context.Context, req *pb.ListAccountRequest) (*pb.ListAccountReply, error) {
	return &pb.ListAccountReply{}, nil
}

func RegisterAccountServer(srv *grpc.Server, service *AccountService) *grpc.Server {
	pb.RegisterAccountServer(srv.Server, service)
	return srv
}

//func RegisterAccountHttpServer(srv *http.Server, service *AccountService) {
//	pb.RegisterAccountServer(srv.Server, service)
//}
