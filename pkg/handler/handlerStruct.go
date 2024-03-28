package handler

import (
	inter "github.com/anjush-bhargavan/go_trade_admin/pkg/service/interfaces"
	pb "github.com/anjush-bhargavan/go_trade_admin/pkg/proto"
)

type AdminHandler struct {
	SVC inter.AdminServiceInter
	pb.AdminServiceServer
}


func NewAdminHandler (svc inter.AdminServiceInter) *AdminHandler {
	return &AdminHandler{
		SVC: svc,
	}
}