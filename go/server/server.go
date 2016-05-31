package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "github.com/dtan4/grpc-pg-todo/go/proto"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type todoServer struct {
	DB *sql.DB
}

const (
	dbHost     = "localhost"
	dbName     = "grpc_pg_todo"
	dbPassword = "password"
	dbPort     = "5432"
	dbUser     = "postgres"
	port       = ":50051"
)

func (s *todoServer) AddTask(ctx context.Context, task *pb.Task) (*pb.Task, error) {
	_, err := s.DB.Exec("INSERT INTO todos(title, description) VALUES($1, $2)", task.Title, task.Description)

	if err != nil {
		return nil, err
	}

	return task, nil
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

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName))

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS todos (title varchar(50), description text)")

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoServer(s, &todoServer{
		DB: db,
	})

	s.Serve(lis)
}
