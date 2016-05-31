package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/dtan4/grpc-pg-todo/go/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	addr  = "localhost:50051"
	usage = `
usage:
  client add <title> <description>
  client list
`
)

func main() {
	var command, title, description string

	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	command = os.Args[1]

	if command == "add" {
		if len(os.Args) != 4 {
			fmt.Println(usage)
			os.Exit(1)
		}

		title = os.Args[2]
		description = os.Args[3]
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewTodoClient(conn)

	switch command {
	case "add":
		task := &pb.Task{
			Title:       title,
			Description: description,
		}

		_, err = client.AddTask(context.Background(), task)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Suceeded.")

	case "list":
		tasks, err := client.ListTasks(context.Background(), &pb.ListRequest{})

		if err != nil {
			log.Fatal(err)
		}

		for _, task := range tasks.Tasks {
			fmt.Printf("%s: %s\n", task.Title, task.Description)
		}

	default:
		fmt.Println(usage)
		os.Exit(1)
	}
}
