package main

import (
	"endProject/pkg/api"
	"endProject/pkg/students"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	a := grpc.NewServer()
	server := &students.GRPCServer{}
	api.RegisterStudentsServer(a, server)
	j, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := a.Serve(j); err != nil {
		log.Fatal(err)
	}
}
