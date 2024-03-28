package di

import (
	"log"

	"github.com/anjush-bhargavan/go_trade_admin/config"
	"github.com/anjush-bhargavan/go_trade_admin/pkg/clients/product"
	"github.com/anjush-bhargavan/go_trade_admin/pkg/clients/user"
	"github.com/anjush-bhargavan/go_trade_admin/pkg/db"
	"github.com/anjush-bhargavan/go_trade_admin/pkg/handler"
	"github.com/anjush-bhargavan/go_trade_admin/pkg/repo"
	"github.com/anjush-bhargavan/go_trade_admin/pkg/server"
	"github.com/anjush-bhargavan/go_trade_admin/pkg/service"
)

// Init initializes the application by setting up its dependencies.
func Init() {
	cnfg := config.LoadConfig()

	db := db.ConnectDB(cnfg)

	productClient, err := product.ClientDial(*cnfg)
	if err != nil {
		log.Fatalf("failed to connect to product client")
	}

	userClient, err := user.ClientDial(*cnfg)
	if err != nil {
		log.Fatalf("failed to connect to user client")
	}

	adminRepo := repo.NewAdminRepository(db)

	adminService := service.NewAdminService(adminRepo,productClient,userClient)

	adminHandler := handler.NewAdminHandler(adminService)

	err = server.NewGrpcAdminServer(cnfg.GrpcPort, adminHandler)
	if err != nil {
		log.Fatalf("failed to start gRPC server %v", err.Error())
	}

}
