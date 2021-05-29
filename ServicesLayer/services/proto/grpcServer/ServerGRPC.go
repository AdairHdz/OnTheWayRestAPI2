package grpcServer

import (
	"context"

	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/proto"
	// "google.golang.org/grpc"
)


const (
	serverURL = "0.0.0.0:4500";
)

type ServerGRPC struct {		
	proto.UnimplementedEmailTokenServer
}

// func New() *ServerGRPC {
// 	return &ServerGRPC{}
// }

func (serverGRPC *ServerGRPC) SendCode(ctx context.Context, request *proto.EmailTokenRequest) (*proto.EmailTokenResponse, error) {
	// conn, err := grpc.Dial(serverURL, grpc.WithInsecure())

	// if err != nil {
	// 	panic("ERROR")
	// }

	// serviceClient := proto.NewEmailTokenClient(conn)

	// serviceClient.GetEmailAndToken(context.Background(), &proto.EmailTokenRequest{
	// 	Email: "adairho16@gmail.com",
	// 	Token: "Holis",
	// })

	response, err := serverGRPC.GetEmailAndToken(context.Background(), request)

	if err != nil {
		panic(err.Error())
	}
	return response, nil
}