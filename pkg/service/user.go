package service


import (
	"context"

	userpb "github.com/anjush-bhargavan/go_trade_admin/pkg/clients/user/pb"
	pb "github.com/anjush-bhargavan/go_trade_admin/pkg/proto"
)

// BlockUserService handle the admin to block the users using the provided information.
func (a *AdminService) BlockUserService(p *pb.AdID) (*pb.AdminResponse, error) {
	ctx := context.Background()

	user := &userpb.ID{
		ID: p.ID,
	}

	_, err := a.UserClient.BlockUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.AdminResponse{
		Status: pb.AdminResponse_OK,
		Message: "User blocked successfully",
	}, nil
}