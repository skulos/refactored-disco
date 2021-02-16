package main

import "intern.mn8.ee/arash/logger"

func main() {

	// conn, _ := grpc.Dial("localhost:12345", grpc.WithInsecure())

	// client := intern.NewGigsafeClient(conn)

	// client.GetEvents(context.Background(), &intern.Empty{})
	// server := intern.GigsafeServer{}
	logger.Setup()

	logger.Sync()
	logger.CleanUp()
}
