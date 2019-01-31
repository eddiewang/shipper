package main

import (
	"fmt"
	"log"

	pb "github.com/eddiewang/shipper/user-service/proto/user"
	"github.com/micro/go-micro"
)

func main() {
	// Create a db connection
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("could not connect to DB: %v", err)
	}

	// MIgrate user  struct
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	s := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("lastest"),
	)

	s.Init()

	pb.RegisterUserServiceHandler(s.Server(), &service{repo, tokenService})

	if err := s.Run(); err != nil {
		fmt.Println(err)
	}

}
