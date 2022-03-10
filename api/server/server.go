package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"app/pb"
	"app/server/db"

	"google.golang.org/grpc"
)

func checkErr(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	db.DbConnect()
}

type server struct {}

func main() {
	fmt.Println("起動")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	checkErr("Failed to listen: %v", err)

	s := grpc.NewServer()
	pb.RegisterAppServiceServer(s, &server{})
	s.Serve(lis)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	fmt.Println("終了中")
	s.Stop()
	lis.Close()
	fmt.Println("終了")
}

func (s *server) GetAllCountry(ctx context.Context, req *pb.AllCountryRequest) (*pb.AllCountryResponse, error) {
	allCountryEntity, err := db.GetAllCountry()
	if err != nil {
		log.Fatalln(err)
	}

	var allCountrySlice []*pb.Country
	for _, countryEntity := range allCountryEntity {
		country := &pb.Country{
			Id:   countryEntity.Id,
			Name: countryEntity.Name,
		}
		allCountrySlice = append(allCountrySlice, country)
	}

	return &pb.AllCountryResponse{AllCountry: allCountrySlice}, nil
}
