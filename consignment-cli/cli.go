package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/eddiewang/shipper/consignment-service/proto/consignment"
	micro "github.com/micro/go-micro"
)

const (
	address         = "localhost:50051"
	defaultFileName = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	service := micro.NewService(micro.Name("go.micro.srv.cli"))
	service.Init()

	client := pb.NewShippingService("go.micro.srv.consignment", service.Client())
	// Contact the server and print out response
	file := defaultFileName
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("created: %t", r.Created)

	// GetConsignments
	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
