package service

import (
	"context"

	productpb "github.com/anjush-bhargavan/go_trade_admin/pkg/clients/product/pb"
	pb "github.com/anjush-bhargavan/go_trade_admin/pkg/proto"
)

// AddCategoryService adds a new category using the provided information.
func (a *AdminService) AddCategoryService(p *pb.AdminCategory) (*pb.AdminResponse, error) {
	ctx := context.Background()

	category := &productpb.Category{
		Name:        p.Name,
		Description: p.Description,
	}

	result, err := a.ProductClient.CreateCategory(ctx, category)
	if err != nil {
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_Status(result.Status),
			Message: result.Message,
			Payload: &pb.AdminResponse_Error{Error: err.Error()},
		}, err
	}

	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "Category created succesfully",
	}, nil
}


// EditCategoryService edits the category using the provided information.
func (a *AdminService) EditCategoryService(p *pb.AdminCategory) (*pb.AdminCategory, error) {
	ctx := context.Background()

	category := &productpb.Category{
		Category_ID: p.Category_ID,
		Name:        p.Name,
		Description: p.Description,
	}

	result, err := a.ProductClient.EditCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	return &pb.AdminCategory{
		Category_ID: result.Category_ID,
		Name:        result.Name,
		Description: result.Description,
	}, nil
}

// RemoveCategoryService removes the category using the provided information.
func (a *AdminService) RemoveCategoryService(p *pb.AdID) (*pb.AdminResponse, error) {
	ctx := context.Background()

	category := &productpb.PrID{
		ID: p.ID,
	}

	_, err := a.ProductClient.RemoveCategory(ctx, category)
	if err != nil {
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_ERROR,
			Message: "error in removing category",
			Payload: &pb.AdminResponse_Error{Error: err.Error()},
		}, err
	}

	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "Category removed succesfully",
	}, nil
}

// FindCategoryService fetch the category using the provided information.
func (a *AdminService) FindCategoryService(p *pb.AdID) (*pb.AdminCategory, error) {
	ctx := context.Background()

	category := &productpb.PrID{
		ID: p.ID,
	}

	result, err := a.ProductClient.FindCategoryByID(ctx, category)
	if err != nil {
		return nil, err
	}

	return &pb.AdminCategory{
		Category_ID: result.Category_ID,
		Name: result.Name,
		Description: result.Description,
	}, nil
}

// FindAllCategoryService finds all categories using the provided information.
func (a *AdminService) FindAllCategoryService(p *pb.AdminNoParam) (*pb.AdminCategoryList, error) {
	ctx := context.Background()

	

	result, err := a.ProductClient.FindAllCategories(ctx,&productpb.ProductNoParam{})
	if err != nil {
		return nil, err
	}

	var categories []*pb.AdminCategory

	for _,category := range result.Categories {
		categories = append(categories, &pb.AdminCategory{
			Category_ID: category.Category_ID,
			Name: category.Name,
			Description: category.Description,
		})
	}

	return &pb.AdminCategoryList{
		Categories: categories,
	}, nil
}