package handler

import (
	"context"

	pb "github.com/anjush-bhargavan/go_trade_admin/pkg/proto"
)

// AddCategory handles the addition of a new category by delegating the operation to the AddCategoryService method of the Admin service.
func (a *AdminHandler) AddCategory(ctx context.Context, p *pb.AdminCategory) (*pb.AdminResponse, error) {
	response, err := a.SVC.AddCategoryService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// FindCategory retrieves a category by its ID by delegating the operation to the FindCategoryService method of the Admin service.
func (a *AdminHandler) FindCategory(ctx context.Context, p *pb.AdID) (*pb.AdminCategory, error) {
	response, err := a.SVC.FindCategoryService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// FindCategories retrieves all categories by delegating the operation to the FindAllCategoryService method of the Admin service.
func (a *AdminHandler) FindCategories(ctx context.Context, p *pb.AdminNoParam) (*pb.AdminCategoryList, error) {
	response, err := a.SVC.FindAllCategoryService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// EditCategoryAdmin handles the editing of an existing category by delegating the operation to the EditCategoryService method of the Admin service.
func (a *AdminHandler) EditCategoryAdmin(ctx context.Context, p *pb.AdminCategory) (*pb.AdminCategory, error) {
	response, err := a.SVC.EditCategoryService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// RemoveCategoryAdmin handles the removal of a category by delegating the operation to the RemoveCategoryService method of the Admin service.
func (a *AdminHandler) RemoveCategoryAdmin(ctx context.Context, p *pb.AdID) (*pb.AdminResponse, error) {
	response, err := a.SVC.RemoveCategoryService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
