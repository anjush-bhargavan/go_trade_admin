package interfaces

import (
	pb "github.com/anjush-bhargavan/go_trade_admin/pkg/proto"
)

type AdminServiceInter interface {
	LoginService(p *pb.AdminLogin) (*pb.AdminResponse,error)
	UpdateWalletService(p *pb.Amount) (*pb.AdminResponse, error)

	AddCategoryService(p *pb.AdminCategory) (*pb.AdminResponse, error)
	EditCategoryService(p *pb.AdminCategory) (*pb.AdminCategory, error)
	RemoveCategoryService(p *pb.AdID) (*pb.AdminResponse, error) 
	FindCategoryService(p *pb.AdID) (*pb.AdminCategory, error)
	FindAllCategoryService(p *pb.AdminNoParam) (*pb.AdminCategoryList, error)

	RemoveProductService(p *pb.AdID) (*pb.AdminResponse, error)
	FindProductService(p *pb.AdID) (*pb.AdminProduct, error)
	FindAllProductService(p *pb.AdminNoParam) (*pb.AdminProductList, error)

	BlockUserService(p *pb.AdID) (*pb.AdminResponse, error)
}