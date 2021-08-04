package main

import (
	"context"
	"fmt"
	v1 "github.com/pupi94/madara/grpc/pb/v1"
	"google.golang.org/grpc"
)

func main() {
	CreateProduct()
}

func CreateProduct() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial("127.0.0.1:3000", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := v1.NewProductControllerClient(conn)
	resp, err := client.CreateProduct(context.Background(), &v1.CreateProductRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println("success = ", resp)
}
