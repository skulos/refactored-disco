package main

import (
	"context"

	"google.golang.org/grpc"
	"intern.mn8.ee/arash/intern"
)

func main() {

	conn, _ := grpc.Dial("localhost:12345", grpc.WithInsecure())

	client := intern.NewGigsafeClient(conn)

	client.GetEvents(context.Background(), &intern.Empty{})
}
