package main

import (
	"fmt"
	"log"

	pb "github.com/dtan4/grpc-pg-todo/go/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	addr = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewTodoClient(conn)

	task := &pb.Task{
		Title:       "dtan4",
		Description: "dtan4 is deprecated",
	}

	_, err = client.AddTask(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Suceeded.")
}
