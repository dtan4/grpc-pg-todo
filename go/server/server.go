package main

import (
	"log"
	"net"

	pb "github.com/dtan4/grpc-pg-todo/go/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type todoServer struct {
}

const (
	port = ":50051"
)

func (s *todoServer) AddTask(ctx context.Context, task *pb.Task) (*pb.Tasks, error) {
	var tasks []*pb.Task

	return &pb.Tasks{
		Tasks: tasks,
	}, nil
}

func (s *todoServer) ListTasks(ctx context.Context, req *pb.ListRequest) (*pb.Tasks, error) {
	var tasks []*pb.Task

	return &pb.Tasks{
		Tasks: tasks,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoServer(s, &todoServer{})
	s.Serve(lis)
}
