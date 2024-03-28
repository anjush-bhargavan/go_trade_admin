package user


import (
	"log"

	"github.com/anjush-bhargavan/go_trade_admin/config"
	pb "github.com/anjush-bhargavan/go_trade_admin/pkg/clients/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.UserServiceClient, error) {
	grpc, err := grpc.Dial(":"+cfg.UserPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc user client: %s, ", cfg.ProductPort)
		return nil, err
	}
	log.Printf("Succesfully connected to user client at port: %v", cfg.ProductPort)
	return pb.NewUserServiceClient(grpc), nil
}