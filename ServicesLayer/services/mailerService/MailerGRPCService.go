package mailerService

import (
	ctx "context"	
	"log"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/proto"
	"google.golang.org/grpc"
)


type MailerGRPCService struct { }

func (MailerGRPCService) SendEmail(emailAddress, token string) (bool) {
	conn, err := grpc.Dial("host.docker.internal:4500", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Println(err.Error())
		return false
	}

	serviceClient := proto.NewEmailTokenClient(conn)			

	response, err := serviceClient.GetEmailAndToken(ctx.Background(), &proto.EmailTokenRequest{
		Email: emailAddress,
		Token: token,
	})

	if err != nil {
		log.Println(err.Error())
		return false
	}	
	return response.Response
}