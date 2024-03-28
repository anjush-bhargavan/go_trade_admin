package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_admin/pkg/proto"
)

// FindProductByID retrieves a product by its ID by delegating the operation to the FindProductService method of the Admin service.
func (a *AdminHandler) FindProductByID(ctx context.Context, p *pb.AdID) (*pb.AdminProduct, error) {
	response, err := a.SVC.FindProductService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// FindAllProducts retrieves all products by delegating the operation to the FindAllProductService method of the Admin service.
func (a *AdminHandler) FindAllProducts(ctx context.Context, p *pb.AdminNoParam) (*pb.AdminProductList, error) {
	response, err := a.SVC.FindAllProductService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// RemoveProduct handles the removal of a product by delegating the operation to the RemoveProductService method of the Admin service.
func (a *AdminHandler) RemoveProduct(ctx context.Context, p *pb.AdID) (*pb.AdminResponse, error) {
	response, err := a.SVC.RemoveProductService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
