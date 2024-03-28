package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_admin/pkg/proto"
)

// AdminBlockUser helps to find the user and block by Admin service.
func (a *AdminHandler) AdminBlockUser(ctx context.Context,p *pb.AdID) (*pb.AdminResponse, error) {
	response, err := a.SVC.BlockUserService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}