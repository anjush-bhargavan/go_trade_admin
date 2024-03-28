package interfaces

import "github.com/anjush-bhargavan/go_trade_admin/pkg/model"

type AdminRepoInter interface {
	FindAdminByEmail(email string) (*model.Admin,error)
	UpdateWallet(amount float64) error 
}