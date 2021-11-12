package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	pb "probufExample/usermgmt"
	"time"
)

const address = "localhost:50051"

func main()  {
	// dial connection
	conn, err :=grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Fatalf("did not connect: %+v", err)
	}
	defer conn.Close()
	c:=pb.NewUserManagementClient(conn)

	//define new context
	ctx, cancel :=context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	// hard code newUser information
	var new_users=make(map[string]int32)
	new_users["Karina"]=22
	new_users["Winter"]=21

	for name, age := range new_users {
		r, err :=c.CreateNewUser(ctx, &pb.NewUser{Name:name, Age: age})
		if err != nil {
			logrus.Fatalf("could not create user: %+v", err)
		}
		log.Printf(`User Details:
Name: %s
Age: %d
ID: %d
		`, r.GetName(), r.GetAge(), r.GetId())
	}
}
