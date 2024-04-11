package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_admin/pkg/proto"
)

func (a *AdminHandler) AddToAdminWallet(ctx context.Context, p *pb.Amount) (*pb.AdminResponse, error) {
	response, err := a.SVC.UpdateWalletService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

