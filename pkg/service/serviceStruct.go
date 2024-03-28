package service

import (
	interRepo "github.com/anjush-bhargavan/go_trade_admin/pkg/repo/interfaces"
	"github.com/anjush-bhargavan/go_trade_admin/pkg/service/interfaces"
	productpb "github.com/anjush-bhargavan/go_trade_admin/pkg/clients/product/pb"
	userpb "github.com/anjush-bhargavan/go_trade_admin/pkg/clients/user/pb"
)
type AdminService struct {
	Repo interRepo.AdminRepoInter
	ProductClient productpb.ProductServiceClient
	UserClient userpb.UserServiceClient
}

func NewAdminService (repo interRepo.AdminRepoInter,productClient productpb.ProductServiceClient,userClient userpb.UserServiceClient) interfaces.AdminServiceInter {
	return &AdminService{
		Repo: repo,
		ProductClient: productClient,
		UserClient: userClient,
	}
}