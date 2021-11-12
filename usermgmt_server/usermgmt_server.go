package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	pb "probufExample/usermgmt"
)

const port = ":50051"

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

// newUser function as we defined in proto file
func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	logrus.Printf("Received: %+v", in.GetName())
	var user_id int32 = int32(rand.Intn(1000))
	return &pb.User{
		Name: in.GetName(),
		Age:  in.GetAge(),
		Id:   user_id,
	}, nil
}

func main()  {
	// listen port
	lis,err:=net.Listen("tcp", port)
	if err!= nil {
		log.Fatalf("failed to listen: %+v", err)
	}

	// generate new grpc server
	s:=grpc.NewServer()
	pb.RegisterUserManagementServer(s,&UserManagementServer{})
	logrus.Printf("server listening at %+v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
}