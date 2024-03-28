package repo

import "github.com/anjush-bhargavan/go_trade_admin/pkg/model"

// FindAdminByEmail method will find the admin data from database using email field
func (a *AdminRepository) FindAdminByEmail(email string) (*model.Admin, error) {
	var admin model.Admin

	if err := a.DB.Model(&model.Admin{}).Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

// UpdateWallet method will update the wallet of admin 
func (a *AdminRepository) UpdateWallet(amount float64) error {
	var admin model.Admin

	if err := a.DB.Model(&model.Admin{}).Where("id = ?", 1).First(&admin).Error; err != nil {
		return err
	}
	admin.Wallet += amount
	return nil
}
