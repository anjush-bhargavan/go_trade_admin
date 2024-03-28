package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_admin/pkg/proto"
)

func (a *AdminHandler) AdminLoginRequest(ctx context.Context, p *pb.AdminLogin) (*pb.AdminResponse, error) {
	response, err := a.SVC.LoginService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
