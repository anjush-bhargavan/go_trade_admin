package repo

import (
	"github.com/anjush-bhargavan/go_trade_admin/pkg/repo/interfaces"
	"gorm.io/gorm"
)

// AdminRepository for connecting db to AdminRepoInter methods
type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepoInter {
	return &AdminRepository{
		DB: db,
	}
}