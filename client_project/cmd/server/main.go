//go:generate make proto
package main

import (
	pb "github.com/abhokarikar/protos/go/myproto/v1"
)

func main() {
	// Create a new instance of the generated protobuf user
	user := &pb.User{
		Id:   "1",
		Name: "Ravi",
	}

	// Print the user
	println("User ID:", user.Id)
	println("User Name:", user.Name)
}
