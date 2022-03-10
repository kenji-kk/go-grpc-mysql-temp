package main

import (
	"app/pb"
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"
)

func checkErr(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func country(w http.ResponseWriter, r *http.Request){
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("server:50051", opts)
	checkErr("could not connect: %v", err)
	defer cc.Close()
	c := pb.NewAppServiceClient(cc)

	res, err := c.GetAllCountry(context.Background(), &pb.AllCountryRequest{})
	checkErr("Fail to create client: %v",err)
	fmt.Printf("AllCountry has geted: %v", res)
}

func main() {
	http.HandleFunc("/", country)
	http.ListenAndServe(":8080", nil)
}
