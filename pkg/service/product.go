package service


import (
	"context"

	productpb "github.com/anjush-bhargavan/go_trade_admin/pkg/clients/product/pb"
	pb "github.com/anjush-bhargavan/go_trade_admin/pkg/proto"
)

// RemoveProductService removes the Product using the provided information.
func (a *AdminService) RemoveProductService(p *pb.AdID) (*pb.AdminResponse, error) {
	ctx := context.Background()

	Product := &productpb.PrIDs{
		Seller_ID: 0,
		ID: p.ID,
	}

	result, err := a.ProductClient.RemoveProduct(ctx, Product)
	if err != nil {
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_Status(result.Status),
			Message: result.Message,
			Payload: &pb.AdminResponse_Error{Error: err.Error()},
		}, err
	}

	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "Product removed succesfully",
	}, nil
}

// FindProductService fetch the Product using the provided information.
func (a *AdminService) FindProductService(p *pb.AdID) (*pb.AdminProduct, error) {
	ctx := context.Background()

	Product := &productpb.PrID{
		ID: p.ID,
	}

	result, err := a.ProductClient.FindProductByID(ctx, Product)
	if err != nil {
		return nil, err
	}

	return &pb.AdminProduct{
		Product_ID: result.Product_ID,
		Name:       result.Name,
		Seller_ID: result.Seller_ID,
		Category: &pb.AdminCategory{
			Category_ID: result.Category.Category_ID,
			Name:        result.Category.Name,
			Description: result.Category.Description,
		},
		Base_Price:         result.Base_Price,
		Bidding_Type:       result.Bidding_Type,
		Bidding_Start_Time: result.Bidding_Start_Time,
		Bidding_End_Time:   result.Bidding_End_Time,
		Listed_On:          result.Listed_On,
		Status:             result.Status,
	}, nil
}

// FindAllProductService finds all products using the provided information.
func (a *AdminService) FindAllProductService(p *pb.AdminNoParam) (*pb.AdminProductList, error) {
	ctx := context.Background()

	

	result, err := a.ProductClient.FindAllProducts(ctx,&productpb.ProductNoParam{})
	if err != nil {
		return nil, err
	}

	var products []*pb.AdminProduct

	for _,product := range result.Products {
		products = append(products, &pb.AdminProduct{
			Product_ID: product.Product_ID,
			Name:       product.Name,
			Seller_ID:  product.Seller_ID,
			Category: &pb.AdminCategory{
				Category_ID: product.Category.Category_ID,
				Name:        product.Category.Name,
				Description: product.Category.Description,
			},
			Base_Price:         product.Base_Price,
			Bidding_Type:       product.Bidding_Type,
			Bidding_Start_Time: product.Bidding_Start_Time,
			Bidding_End_Time:   product.Bidding_End_Time,
			Listed_On:          product.Listed_On,
			Status:             product.Status,
		})
	}

	return &pb.AdminProductList{
		Products: products,
	}, nil
}