package main

import (
	"context"
	"database/sql"
	"endProject/pkg/api"
	"endProject/pkg/students"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test"
	"log"
	"net"
	"net/http"
)

func main() {
	go func() {
		//mux
		mux := runtime.NewServeMux()
		pb.RegisterTestApiHandlerServer(context.Background(), mux, &test.ApiServer{})
		log.Fatal(http.ListenAndServe("localhost:8081", mux))
	}()
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
	connStr := "user= postgres password=postgres dbname= postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
