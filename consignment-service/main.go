package main

import (
	"log"
	"os"

	pb "github.com/eddiewang/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/eddiewang/shipper/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	// Db Host
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	if err != nil {
		log.Panicf("could not connect to datastore with host %s - %v", host, err)
	}

	defer session.Close()

	srv := micro.NewService(
		// name must match package name in protobuf
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselService("go.micro.srv.vessel", srv.Client())

	srv.Init()
	// Register service with gRPC server, which ties our implementation into the
	// auto-generated interface code for our protobuf definition.
	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
