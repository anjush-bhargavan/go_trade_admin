package service

import (
	"fmt"

	pb "github.com/anjush-bhargavan/go_trade_admin/pkg/proto"
)

func (a *AdminService) UpdateWalletService(p *pb.Amount) (*pb.AdminResponse, error) {

	wallet,err := a.Repo.UpdateWallet(float64(p.Amount))
	if err != nil {
		return &pb.AdminResponse{
			Status: pb.AdminResponse_ERROR,
			Message: "error while updating amount to wallet",
			Payload: &pb.AdminResponse_Error{Error: err.Error()},
		},err
	}


	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "Wallet updated successful",
		Payload: &pb.AdminResponse_Data{Data:fmt.Sprintf("wallet updated, new wallet amount is : %f",wallet) },
	}, nil
}
