package product


import (
	"log"

	"github.com/anjush-bhargavan/go_trade_admin/config"
	pb "github.com/anjush-bhargavan/go_trade_admin/pkg/clients/product/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.ProductServiceClient, error) {
	grpc, err := grpc.Dial(":"+cfg.ProductPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc product client: %s, ", cfg.ProductPort)
		return nil, err
	}
	log.Printf("Succesfully connected to product client at port: %v", cfg.ProductPort)
	return pb.NewProductServiceClient(grpc), nil
}