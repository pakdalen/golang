package main

import (
	"context"
	"endProject/pkg/api"
	"flag"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {
	flag.Parse()
	if flag.NArg() < 6 {
		log.Fatal("add more arguments")
	}
	Name := flag.Arg(0)

	Surname := flag.Arg(1)

	Id, err := strconv.Atoi(flag.Arg(2))
	if err != nil {
		log.Fatal(err)
	}
	Yearofstudy, err := strconv.Atoi(flag.Arg(3))
	if err != nil {
		log.Fatal(err)
	}
	Marks, err := strconv.Atoi(flag.Arg(4))
	if err != nil {
		log.Fatal(err)
	}
	Subjects, err := strconv.Atoi(flag.Arg(5))
	if err != nil {
		log.Fatal(err)
	}
	connection, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := api.NewStudentsClient(connection)
	res, err := c.GetStudent(context.Background(), &api.GetStudentListRequest{Name: string(Name), Surname: string(Surname), Id: int32(Id), Yearofstudy: int32(Yearofstudy), Marks: int32(Marks), Subjects: int32(Subjects)})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.GetGpa())

}
